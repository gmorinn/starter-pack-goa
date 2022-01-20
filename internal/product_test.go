package db

import (
	"api_crud/utils"
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	arg := CreateProductParams{
		Name:     utils.RandStringRunes(12),
		Price:    float64(utils.RandomInt(0, 299)),
		Cover:    utils.RandStringRunes(50),
		Category: Categories(utils.RandomAttribut([]string{"sneaker", "hat", "women", "men", "jacket"})),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.NotEmpty(t, product.Category)
	require.NotEmpty(t, product.Cover)
	require.NotEmpty(t, product.ID)
	require.NotEmpty(t, product.Name)
	require.NotEmpty(t, product.Price)
	require.NotEmpty(t, product.CreatedAt)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Category, product.Category)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Cover, product.Cover)
	require.Equal(t, product.CreatedAt, product.UpdatedAt)

	require.Positive(t, product.Price)
	require.Empty(t, product.DeletedAt)
}
