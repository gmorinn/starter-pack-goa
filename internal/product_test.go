package db

import (
	"api_crud/utils"
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:     utils.RandStringRunes(12),
		Price:    float64(utils.RandomInt(1, 299)),
		Cover:    utils.RandStringRunes(50),
		Category: Categories(utils.RandomAttribut([]string{"sneaker", "hat", "women", "men", "jacket"})),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.NotEmpty(t, product.ID)
	require.GreaterOrEqual(t, product.Price, 0.0)
	require.NotEmpty(t, product.CreatedAt)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Category, product.Category)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Cover, product.Cover)
	require.Equal(t, product.CreatedAt, product.UpdatedAt)

	require.Positive(t, product.Price)
	require.Empty(t, product.DeletedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	newProduct := createRandomProduct(t)
	product, err := testQueries.GetProduct(context.Background(), newProduct.ID)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, newProduct.ID, product.ID)
	require.Equal(t, newProduct.Cover, product.Cover)
	require.Equal(t, newProduct.Name, product.Name)
	require.Equal(t, newProduct.Category, product.Category)
	require.Equal(t, newProduct.Price, product.Price)
}

func TestUpdateProduct(t *testing.T) {
	newProduct := createRandomProduct(t)

	arg := UpdateProductParams{
		ID:       newProduct.ID,
		Name:     newProduct.Name,
		Price:    float64(utils.RandomInt(0, 299)),
		Cover:    newProduct.Cover,
		Category: Categories(utils.RandomAttribut([]string{"sneaker", "hat", "women", "men", "jacket"})),
	}

	err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)

	product, err := testQueries.GetProduct(context.Background(), newProduct.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, newProduct.ID, product.ID)
	require.Equal(t, newProduct.Cover, product.Cover)
	require.Equal(t, newProduct.Name, product.Name)
	require.Equal(t, arg.Category, product.Category)
	require.Equal(t, arg.Price, arg.Price)
}

func TestDeleteProduct(t *testing.T) {
	newProduct := createRandomProduct(t)

	err := testQueries.DeleteProduct(context.Background(), newProduct.ID)
	require.NoError(t, err)

	product, err := testQueries.GetProduct(context.Background(), newProduct.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product)
}

func TestListBoProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := GetBoAllProductsParams{
		NameAsc: true,
		Limit:   5,
		Offset:  5,
	}

	products, err := testQueries.GetBoAllProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, v := range products {
		require.NotEmpty(t, v)
	}
}

func TestCountProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	var min int64 = 9
	count, err := testQueries.GetCountsProducts(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, min)
}

func TestGetProductsByCategory(t *testing.T) {
	arg := CreateProductParams{
		Name:     utils.RandStringRunes(12),
		Price:    float64(utils.RandomInt(1, 299)),
		Cover:    utils.RandStringRunes(50),
		Category: CategoriesHat,
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	arg = CreateProductParams{
		Name:     utils.RandStringRunes(12),
		Price:    float64(utils.RandomInt(1, 299)),
		Cover:    utils.RandStringRunes(50),
		Category: CategoriesHat,
	}
	product, err = testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	products, err := testQueries.GetProductsByCategory(context.Background(), CategoriesHat)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(products), 2)

	for _, v := range products {
		require.NotEmpty(t, v)
		require.Equal(t, v.Category, CategoriesHat)
	}
}

func TestGetAllProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	products, err := testQueries.GetAllProducts(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(products), 9)

	for _, v := range products {
		require.NotEmpty(t, v)
	}
}
