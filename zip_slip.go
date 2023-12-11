package unzip

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsZipSlip 对单个文件进行zip slip检查
// baseDirectory: 要解压到的目录
// filename: 要解压的文件的名称
func IsZipSlip(baseDirectory, filename string) bool {
	if baseDirectory == "" {
		baseDirectory = "fake-directory"
	}
	fileUnzipPath := filepath.Clean(filepath.Join(baseDirectory, filename))
	fileUnzipDirectory := filepath.Clean(baseDirectory) + string(os.PathSeparator)
	return !strings.HasPrefix(fileUnzipPath, fileUnzipDirectory)
}

// BuildZipSlipError 当检查到zip slip文件时，为其构造错误信息
func BuildZipSlipError(file *zip.File) error {
	// TODO 完善错误信息
	return fmt.Errorf("find zip slip %s", file.Name)
}
