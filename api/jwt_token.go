package api

import (
	jwttoken "api_crud/gen/jwt_token"
	db "api_crud/internal"
	"api_crud/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
	"gopkg.in/oauth2.v3/utils/uuid"
)

type jwtTokensrvc struct {
	logger *log.Logger
	server *Server
}

func errorEmail() *jwttoken.EmailAlreadyExist {
	return &jwttoken.EmailAlreadyExist{
		Message: "EMAIL_ALREADY_EXIST",
	}
}

func NewJWTToken(logger *log.Logger, server *Server) jwttoken.Service {
	return &jwtTokensrvc{logger, server}
}

func (s *jwtTokensrvc) errorResponse(msg string, err error) *jwttoken.UnknownError {
	return &jwttoken.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

var (
	ErrInvalidToken       error = jwttoken.Unauthorized("invalid token")
	ErrInvalidPassword    error = jwttoken.Unauthorized("invalid password")
	ErrInvalidTokenScopes error = jwttoken.InvalidScopes("invalid scopes in token")
	ErrExpiredToken       error = jwttoken.Unauthorized("token has expired")
)

func (s *jwtTokensrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

func (s *jwtTokensrvc) Signup(ctx context.Context, p *jwttoken.SignupPayload) (res *jwttoken.Sign, err error) {

	if p.Password != p.ConfirmPassword {
		return nil, ErrInvalidPassword
	}

	isExist, err := s.server.Store.ExistUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, s.errorResponse("ERROR_GET_MAIL", err)
	}
	if isExist {
		return nil, errorEmail()
	}

	arg := db.SignupParams{
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Email:     p.Email,
		Crypt:     p.Password,
		Phone:     utils.NullS(p.Phone),
		Birthday:  utils.NullS(p.Birthday),
	}
	user, err := s.server.Store.Signup(ctx, arg)
	if err != nil {
		return nil, s.errorResponse("ERROR_CREATE_USER", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(user.ID), string(user.Role))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) Signin(ctx context.Context, p *jwttoken.SigninPayload) (res *jwttoken.Sign, err error) {
	// Request Login
	arg := db.LoginUserParams{
		Email: p.Email,
		Crypt: p.Password,
	}
	user, err := s.server.Store.LoginUser(ctx, arg)
	if err != nil {
		return nil, s.errorResponse("ERROR_LOGIN_USER", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(user.ID), string(user.Role))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) Refresh(ctx context.Context, p *jwttoken.RefreshPayload) (res *jwttoken.Sign, err error) {
	token, err := jwt.Parse(p.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})

	if err != nil {
		return nil, s.errorResponse("TOKEN_ERROR", err)
	}

	if !token.Valid {
		return nil, s.errorResponse("TOKEN_IS_NOT_VALID", err)
	}

	refresh, err := s.server.Store.GetRefreshToken(ctx, p.RefreshToken)
	if err != nil {
		return nil, s.errorResponse("FIND_REFRESH_TOKEN", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(refresh.UserID), string(refresh.UserRole))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, refresh.UserID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	response := jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return &response, nil
}

func (s *jwtTokensrvc) generateJwtToken(ID uuid.UUID, role string) (string, string, time.Time, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"role":   role,
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	t, err := accessToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_ACCESS_JWT %v", err)
	}

	expt := time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(s.server.Config.Security.RefreshTokenDuration)))
	exp := expt.Unix()

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"role":   role,
		"exp":    exp,
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(s.server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_REFRESH_JWT %v", err)
	}

	return t, r, expt, nil
}

func (s *jwtTokensrvc) AuthProviders(ctx context.Context, p *jwttoken.AuthProvidersPayload) (res *jwttoken.Sign, err error) {
	var t, r string
	var user db.User
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		app, err := firebaseClient(ctx)
		if err != nil {
			return fmt.Errorf("FIREBASE_CLIENT %v", err.Error())
		}

		client, err := app.Auth(ctx)
		if err != nil {
			return fmt.Errorf("FIREBASE_AUTH %v", err.Error())
		}

		_, err = client.VerifyIDTokenAndCheckRevoked(ctx, p.FirebaseIDToken)
		if err != nil {
			return fmt.Errorf("VERIFYING_ID_TOKEN: %v", err.Error())
		}
		// CHECK IF USER BY FIREBASE EXIST
		existFBUid, err := q.ExistGetUserByFireBaseUid(ctx, utils.NullS(p.FirebaseUID))
		if err != nil {
			return fmt.Errorf("EXIST_GET_USER_BY_FIRE_BASE_UID %v", err.Error())
		}
		// IF USER WAS EVER CONNECTED WITH FIREBASE
		if existFBUid {
			user, err = q.GetUserByFireBaseUid(ctx, utils.NullS(p.FirebaseUID))
			if err != nil {
				return fmt.Errorf("GET_USER_BY_FIRE_BASE_UID %v", err.Error())
			}
		} else {
			// CHECK IF USER ALREADY EXIST IN DATABASE
			existEmail, err := q.ExistUserByEmail(ctx, p.Email)
			if err != nil {
				return fmt.Errorf("EXIST_EMAIL: %v", err.Error())
			}
			// IF ALREADY EXIST
			if existEmail {
				user, err = q.FindUserByEmail(ctx, p.Email)
				if err != nil {
					return fmt.Errorf("FIND_USER_BY_EMAIL: %v", err.Error())
				}
				// UPDATE FIREBASE FIELDS IN DB
				if err := q.UpdateUserProvider(ctx, db.UpdateUserProviderParams{
					ID:               user.ID,
					FirebaseIDToken:  utils.NullS(p.FirebaseIDToken),
					FirebaseUid:      utils.NullS(p.FirebaseUID),
					FirebaseProvider: utils.NullS(p.FirebaseProvider)}); err != nil {
					return fmt.Errorf("UPDATE_USER_PROVIDER: %v", err.Error())
				}
			} else {
				arg := db.SignProviderParams{
					Firstname:        p.Firstname,
					Lastname:         p.Lastname,
					Email:            p.Email,
					Crypt:            utils.RandStringRunes(60),
					FirebaseIDToken:  utils.NullS(p.FirebaseIDToken),
					FirebaseUid:      utils.NullS(p.FirebaseUID),
					FirebaseProvider: utils.NullS(p.FirebaseProvider),
				}
				//Sign with Provider
				user, err = q.SignProvider(ctx, arg)
				if err != nil {
					return fmt.Errorf("SIGNUP_PROVIDER %v", err.Error())
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, s.errorResponse("TX_AUTH_PROVIDER", err)
	}

	t, r, expt, err := s.generateJwtToken(uuid.UUID(user.ID), string(user.Role))
	if err != nil {
		return nil, s.errorResponse("ERROR_TOKEN", err)
	}

	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, s.errorResponse("ERROR_REFRESH_TOKEN", err)
	}

	res = &jwttoken.Sign{
		AccessToken:  t,
		RefreshToken: r,
		Success:      true,
	}
	return res, nil
}

func (server *Server) CheckJWT(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["id"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["exp"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := schema.Validate(scopesInToken); err != nil {
		return ctx, jwttoken.InvalidScopes(err.Error())
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		jwtToken: claims,
	})
	return ctx, nil
}

func (s *jwtTokensrvc) EmailExist(ctx context.Context, p *jwttoken.EmailExistPayload) (res *jwttoken.EmailExistResult, err error) {
	var isExist bool = false
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		email, err := q.CheckEmailExist(ctx, p.Email)
		if err != nil {
			return fmt.Errorf("ERROR_GET_EMAIL_EXIST_BY_ID %v", err)
		}
		if email {
			isExist = true
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_EMAIL_EXIST", err)
	}
	if !isExist {
		return &jwttoken.EmailExistResult{
			Success: true,
			Exist:   false,
		}, nil
	}
	return &jwttoken.EmailExistResult{
		Success: true,
		Exist:   true,
	}, nil
}
