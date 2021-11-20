package api

import (
	products "api_crud/gen/products"
	db "api_crud/internal"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// products service example implementation.
// The example methods log the requests and return zero values.
type productssrvc struct {
	logger *log.Logger
	server *Server
}

// NewProducts returns the products service implementation.
func NewProducts(logger *log.Logger, server *Server) products.Service {
	return &productssrvc{logger, server}
}

func (s *productssrvc) errorResponse(msg string, err error) *products.UnknownError {
	return &products.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "products" for the
// "OAuth2" security scheme.
func (s *productssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "products" for the
// "jwt" security scheme.
func (s *productssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Get All products by category
func (s *productssrvc) GetAllProductsByCategory(ctx context.Context, p *products.GetAllProductsByCategoryPayload) (res *products.GetAllProductsByCategoryResult, err error) {
	var ProductsResponse []*products.ResProduct
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		pS, err := s.server.Store.GetProductsByCategory(ctx, db.Categories(p.Category))
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_PRODUCTS %v", err)
		}
		for _, v := range pS {
			id := v.ID.String()
			ProductsResponse = append(ProductsResponse, &products.ResProduct{
				ID:       id,
				Name:     v.Name,
				Price:    v.Price,
				Cover:    v.Cover,
				Category: string(v.Category),
			})
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_PRODUCTS", err)
	}
	res = &products.GetAllProductsByCategoryResult{
		Products: ProductsResponse,
		Success:  true,
	}
	return res, nil
}

// Get one product
func (s *productssrvc) GetProduct(ctx context.Context, p *products.GetProductPayload) (res *products.GetProductResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		b, err := q.GetProduct(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_PRODUCT_BY_ID %v", err)
		}
		res = &products.GetProductResult{
			Product: &products.ResProduct{
				ID:       b.ID.String(),
				Name:     b.Name,
				Cover:    b.Cover,
				Price:    b.Price,
				Category: string(b.Category),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_PRODUCT_ID", err)
	}
	return res, nil
}
