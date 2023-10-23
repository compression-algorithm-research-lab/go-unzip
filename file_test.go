package unzip

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFile_SafeUnzip(t *testing.T) {
	options := NewOptions().SetSourceZipFile("test_data/foo.zip")
	slice, err := New(options).Files()
	assert.Nil(t, err)
	assert.True(t, len(slice) > 0)
	for _, file := range slice {
		err = file.SafeUnzip("output")
		assert.Nil(t, err)
	}

}

func TestFile_ReadBytes(t *testing.T) {
	options := NewOptions().SetSourceZipFile("test_data/foo.zip")
	slice, err := New(options).Files()
	assert.Nil(t, err)
	assert.True(t, len(slice) > 0)

	bytes, err := slice[1].ReadBytes()
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

}

func TestFile_Save(t *testing.T) {
	options := NewOptions().SetSourceZipFile("test_data/foo.zip")
	slice, err := New(options).Files()
	assert.Nil(t, err)
	assert.True(t, len(slice) > 0)

	unzipFilepath := "test_data/unzip-file.txt"
	err = slice[1].Save(unzipFilepath)
	assert.Nil(t, err)

	bytes, err := os.ReadFile(unzipFilepath)
	assert.Nil(t, err)
	assert.Equal(t, "barbar", string(bytes))

}
