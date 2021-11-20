package api

import (
	boproducts "api_crud/gen/bo_products"
	db "api_crud/internal"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// boProducts service example implementation.
// The example methods log the requests and return zero values.
type boProductssrvc struct {
	logger *log.Logger
	server *Server
}

// NewBoProducts returns the boProducts service implementation.
func NewBoProducts(logger *log.Logger, server *Server) boproducts.Service {
	return &boProductssrvc{logger, server}
}

func (s *boProductssrvc) errorResponse(msg string, err error) *boproducts.UnknownError {
	return &boproducts.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "boProducts" for
// the "OAuth2" security scheme.
func (s *boProductssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "boProducts" for the
// "jwt" security scheme.
func (s *boProductssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Get All products
func (s *boProductssrvc) GetAllProducts(ctx context.Context, p *boproducts.GetAllProductsPayload) (res *boproducts.GetAllProductsResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		p, err := s.server.Store.GetAllProducts(ctx)
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_PRODUCTS %v", err)
		}
		var ProductsResponse []*boproducts.ResBoProduct
		for _, v := range p {
			id := v.ID.String()
			ProductsResponse = append(ProductsResponse, &boproducts.ResBoProduct{
				ID:       id,
				Name:     v.Name,
				Price:    v.Price,
				Cover:    v.Cover,
				Category: string(v.Category),
			})

			res = &boproducts.GetAllProductsResult{
				Products: ProductsResponse,
				Success:  true,
			}
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_PRODUCTS", err)
	}
	return res, nil
}

// Get All products by category
func (s *boProductssrvc) GetAllProductsByCategory(ctx context.Context, p *boproducts.GetAllProductsByCategoryPayload) (res *boproducts.GetAllProductsByCategoryResult, err error) {
	var ProductsResponse []*boproducts.ResBoProduct
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		pS, err := s.server.Store.GetProductsByCategory(ctx, db.Categories(p.Category))
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_PRODUCTS %v", err)
		}
		for _, v := range pS {
			id := v.ID.String()
			ProductsResponse = append(ProductsResponse, &boproducts.ResBoProduct{
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
	res = &boproducts.GetAllProductsByCategoryResult{
		Products: ProductsResponse,
		Success:  true,
	}
	return res, nil
}

// Delete one product by ID
func (s *boProductssrvc) DeleteProduct(ctx context.Context, p *boproducts.DeleteProductPayload) (res *boproducts.DeleteProductResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteProduct(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("ERROR_DELETE_PRODUCT_BY_ID %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_PRODUCT", err)
	}
	return &boproducts.DeleteProductResult{Success: true}, nil
}

// Create one product
func (s *boProductssrvc) CreateProduct(ctx context.Context, p *boproducts.CreateProductPayload) (res *boproducts.CreateProductResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateProductParams{
			Name:     p.Product.Name,
			Price:    p.Product.Price,
			Cover:    p.Product.Cover,
			Category: db.Categories(p.Product.Category),
		}
		createdProduct, err := q.CreateProduct(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_PRODUCT %v", err)
		}
		newProduct, err := q.GetProduct(ctx, createdProduct.ID)
		if err != nil {
			return fmt.Errorf("ERROR_GET_PRODUCT_BY_ID %v", err)
		}
		res = &boproducts.CreateProductResult{
			Product: &boproducts.ResBoProduct{
				ID:       newProduct.ID.String(),
				Name:     newProduct.Name,
				Cover:    newProduct.Cover,
				Price:    newProduct.Price,
				Category: string(newProduct.Category),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_PRODUCT", err)
	}
	return res, nil
}

// Update one product
func (s *boProductssrvc) UpdateProduct(ctx context.Context, p *boproducts.UpdateProductPayload) (res *boproducts.UpdateProductResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateProductParams{
			ID:       uuid.MustParse(p.ID),
			Name:     p.Product.Name,
			Price:    p.Product.Price,
			Cover:    p.Product.Cover,
			Category: db.Categories(p.Product.Category),
		}
		if err := q.UpdateProduct(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_PRODUCT %v", err)
		}
		newProduct, err := q.GetProduct(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_PRODUCT_BY_ID %v", err)
		}
		res = &boproducts.UpdateProductResult{
			Product: &boproducts.ResBoProduct{
				ID:       newProduct.ID.String(),
				Name:     newProduct.Name,
				Cover:    newProduct.Cover,
				Price:    newProduct.Price,
				Category: string(newProduct.Category),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_PRODUCT", err)
	}
	return res, nil
}

// Get one product
func (s *boProductssrvc) GetProduct(ctx context.Context, p *boproducts.GetProductPayload) (res *boproducts.GetProductResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		b, err := q.GetProduct(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_PRODUCT_BY_ID %v", err)
		}
		res = &boproducts.GetProductResult{
			Product: &boproducts.ResBoProduct{
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

func (s *boProductssrvc) DeleteManyProducts(ctx context.Context, p *boproducts.DeleteManyProductsPayload) (res *boproducts.DeleteManyProductsResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		for _, v := range p.Tab {
			if err := q.DeleteProduct(ctx, uuid.MustParse(v)); err != nil {
				return fmt.Errorf("ERROR_DELETE_PRODUCT_BY_ID_%v %v", v, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_PRODUCTS", err)
	}
	return &boproducts.DeleteManyProductsResult{Success: true}, nil
}
