package db

import (
	"api_crud/utils"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomFile(t *testing.T) CreateFileRow {
	upldir, _ := utils.UploadDir()
	name := utils.RandStringRunes(10)
	ext := filepath.Ext(utils.RandStringRunes(10))
	un := fmt.Sprintf("%s%s", uuid.New(), ext)
	fn := fmt.Sprintf("%s/%s", upldir, un)

	arg := CreateFileParams{
		Name: utils.NullS(name),
		Url:  utils.NullS("/" + fn),
		Mime: utils.NullS("i"),
		Size: utils.NullI64(utils.RandomInt(1, 900)),
	}

	file, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)
	require.NotEmpty(t, file.Mime)
	require.Equal(t, arg.Name, file.Name)
	require.Equal(t, arg.Url, file.Url)

	return file
}

func TestCreateFile(t *testing.T) {
	createRandomFile(t)
}

func TestDeleteFile(t *testing.T) {
	file := createRandomFile(t)

	err := testQueries.DeleteFile(context.Background(), file.Url)
	require.NoError(t, err)
}

func TestGetFile(t *testing.T) {
	// Good url
	file := createRandomFile(t)
	getFile, err := testQueries.GetFileByURL(context.Background(), file.Url)
	require.NoError(t, err)
	require.NotEmpty(t, getFile)
	require.Equal(t, file.Url, getFile.Url)
	require.Equal(t, file.Name.String, getFile.Name.String)

	// Wrong url
	getFile, err = testQueries.GetFileByURL(context.Background(), utils.NullS("0"))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
