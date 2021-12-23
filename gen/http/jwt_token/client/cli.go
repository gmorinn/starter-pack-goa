// Code generated by goa v3.5.2, DO NOT EDIT.
//
// jwtToken HTTP client CLI support package
//
// Command:
// $ goa gen api_crud/design

package client

import (
	jwttoken "api_crud/gen/jwt_token"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildSignupPayload builds the payload for the jwtToken signup endpoint from
// CLI flags.
func BuildSignupPayload(jwtTokenSignupBody string, jwtTokenSignupOauth string) (*jwttoken.SignupPayload, error) {
	var err error
	var body SignupRequestBody
	{
		err = json.Unmarshal([]byte(jwtTokenSignupBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"birthday\": \"Quibusdam et voluptatem atque blanditiis.\",\n      \"confirm_password\": \"JeSuisUnTest974\",\n      \"email\": \"guillaume@epitech.eu\",\n      \"firstname\": \"Guillaume\",\n      \"lastname\": \"Morin\",\n      \"password\": \"JeSuisUnTest974\",\n      \"phone\": \"+262 692 12 34 56\"\n   }'")
		}
		if utf8.RuneCountInString(body.Firstname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", body.Firstname, utf8.RuneCountInString(body.Firstname), 3, true))
		}
		if utf8.RuneCountInString(body.Firstname) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", body.Firstname, utf8.RuneCountInString(body.Firstname), 15, false))
		}
		if utf8.RuneCountInString(body.Lastname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", body.Lastname, utf8.RuneCountInString(body.Lastname), 3, true))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", body.Password, "\\d"))
		if utf8.RuneCountInString(body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 8, true))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.confirm_password", body.ConfirmPassword, "\\d"))
		if utf8.RuneCountInString(body.ConfirmPassword) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.confirm_password", body.ConfirmPassword, utf8.RuneCountInString(body.ConfirmPassword), 8, true))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if jwtTokenSignupOauth != "" {
			oauth = &jwtTokenSignupOauth
		}
	}
	v := &jwttoken.SignupPayload{
		Firstname:       body.Firstname,
		Lastname:        body.Lastname,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
		Email:           body.Email,
		Birthday:        body.Birthday,
		Phone:           body.Phone,
	}
	{
		var zero string
		if v.Birthday == zero {
			v.Birthday = ""
		}
	}
	{
		var zero string
		if v.Phone == zero {
			v.Phone = ""
		}
	}
	v.Oauth = oauth

	return v, nil
}

// BuildSigninPayload builds the payload for the jwtToken signin endpoint from
// CLI flags.
func BuildSigninPayload(jwtTokenSigninBody string, jwtTokenSigninOauth string) (*jwttoken.SigninPayload, error) {
	var err error
	var body SigninRequestBody
	{
		err = json.Unmarshal([]byte(jwtTokenSigninBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"guillaume@epitech.eu\",\n      \"password\": \"JeSuisUnTest974\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", body.Password, "\\d"))
		if utf8.RuneCountInString(body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 8, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if jwtTokenSigninOauth != "" {
			oauth = &jwtTokenSigninOauth
		}
	}
	v := &jwttoken.SigninPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	v.Oauth = oauth

	return v, nil
}

// BuildSigninBoPayload builds the payload for the jwtToken signin Bo endpoint
// from CLI flags.
func BuildSigninBoPayload(jwtTokenSigninBoBody string, jwtTokenSigninBoOauth string) (*jwttoken.SigninBoPayload, error) {
	var err error
	var body SigninBoRequestBody
	{
		err = json.Unmarshal([]byte(jwtTokenSigninBoBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"guillaume@epitech.eu\",\n      \"password\": \"JeSuisUnTest974\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", body.Password, "\\d"))
		if utf8.RuneCountInString(body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 8, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if jwtTokenSigninBoOauth != "" {
			oauth = &jwtTokenSigninBoOauth
		}
	}
	v := &jwttoken.SigninBoPayload{
		Email:    body.Email,
		Password: body.Password,
	}
	v.Oauth = oauth

	return v, nil
}

// BuildRefreshPayload builds the payload for the jwtToken refresh endpoint
// from CLI flags.
func BuildRefreshPayload(jwtTokenRefreshBody string, jwtTokenRefreshOauth string) (*jwttoken.RefreshPayload, error) {
	var err error
	var body RefreshRequestBody
	{
		err = json.Unmarshal([]byte(jwtTokenRefreshBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"refresh_token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ\"\n   }'")
		}
	}
	var oauth *string
	{
		if jwtTokenRefreshOauth != "" {
			oauth = &jwtTokenRefreshOauth
		}
	}
	v := &jwttoken.RefreshPayload{
		RefreshToken: body.RefreshToken,
	}
	v.Oauth = oauth

	return v, nil
}

// BuildAuthProvidersPayload builds the payload for the jwtToken auth-providers
// endpoint from CLI flags.
func BuildAuthProvidersPayload(jwtTokenAuthProvidersBody string, jwtTokenAuthProvidersOauth string) (*jwttoken.AuthProvidersPayload, error) {
	var err error
	var body AuthProvidersRequestBody
	{
		err = json.Unmarshal([]byte(jwtTokenAuthProvidersBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"guillaume@epitech.eu\",\n      \"firebase_id_token\": \"xrj\",\n      \"firebase_provider\": \"facebook.com\",\n      \"firebase_uid\": \"zgmURRUlcJfgDMRyjJ20xs7Rxxw2\",\n      \"firstname\": \"Guillaume\",\n      \"lastname\": \"Morin\"\n   }'")
		}
		if utf8.RuneCountInString(body.Firstname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", body.Firstname, utf8.RuneCountInString(body.Firstname), 3, true))
		}
		if utf8.RuneCountInString(body.Firstname) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", body.Firstname, utf8.RuneCountInString(body.Firstname), 15, false))
		}
		if utf8.RuneCountInString(body.Lastname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", body.Lastname, utf8.RuneCountInString(body.Lastname), 3, true))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if utf8.RuneCountInString(body.FirebaseIDToken) < 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firebase_id_token", body.FirebaseIDToken, utf8.RuneCountInString(body.FirebaseIDToken), 400, true))
		}
		if utf8.RuneCountInString(body.FirebaseUID) < 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firebase_uid", body.FirebaseUID, utf8.RuneCountInString(body.FirebaseUID), 15, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if jwtTokenAuthProvidersOauth != "" {
			oauth = &jwtTokenAuthProvidersOauth
		}
	}
	v := &jwttoken.AuthProvidersPayload{
		Firstname:        body.Firstname,
		Lastname:         body.Lastname,
		Email:            body.Email,
		FirebaseIDToken:  body.FirebaseIDToken,
		FirebaseUID:      body.FirebaseUID,
		FirebaseProvider: body.FirebaseProvider,
	}
	v.Oauth = oauth

	return v, nil
}
