package unzip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testZipPath = "./test_data/foo.zip"

func TestNew(t *testing.T) {
	unzip := New(NewOptions().SetSourceZipFile(testZipPath))
	assert.NotNil(t, unzip)
}

func TestUnzip_SafeTraversal(t *testing.T) {
	unzip := New(NewOptions().SetSourceZipFile(testZipPath))
	assert.NotNil(t, unzip)

	err := unzip.SafeTraversal(func(file *File, options *Options) error {
		t.Log(file.Name)
		return nil
	})
	assert.Nil(t, err)
}

func TestUnzip_Traversal(t *testing.T) {
	unzip := New(NewOptions().SetSourceZipFile(testZipPath))
	assert.NotNil(t, unzip)

	err := unzip.Traversal(func(file *File, options *Options) error {
		t.Log(file.Name)
		return nil
	})
	assert.Nil(t, err)
}

func TestUnzip_Unzip(t *testing.T) {
	unzip := New(NewOptions().SetSourceZipFile(testZipPath).SetDestinationDirectory("./test_data/a/"))
	assert.NotNil(t, unzip)

	err := unzip.Unzip()
	assert.Nil(t, err)
}

func TestUnzip_Files(t *testing.T) {
	unzip := New(NewOptions().SetSourceZipFile(testZipPath))
	assert.NotNil(t, unzip)

	fileSlice, err := unzip.Files()
	assert.Nil(t, err)
	assert.True(t, len(fileSlice) == 2)
}
