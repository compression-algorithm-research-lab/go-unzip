package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	*zip.File
}

// ReadBytes 读取此文件的字节
func (x *File) ReadBytes() (bytes []byte, err error) {
	open, err := x.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = open.Close()
	}()
	return io.ReadAll(open)
}

// Save 把此文件保存到硬盘上给定的位置，不会自动创建父目录
func (x *File) Save(path string) (err error) {
	if x.FileInfo().IsDir() {
		// 如果是目录的话，则创建对应的目录
		return os.MkdirAll(path, x.Mode())
	} else {
		// 如果是文件的话，则读取文件写入到给定路径
		bytes, err := x.ReadBytes()
		if err != nil {
			return err
		}
		// TODO 这个x.Mode()会不会因为某些特殊的mode而导致写入失败呢
		return os.WriteFile(path, bytes, x.Mode())
	}
}

// Unzip 把文件解压到给定的路径
func (x *File) Unzip(baseDirectory string) (err error) {
	saveToPath := filepath.Join(baseDirectory, x.Name)
	return x.Save(saveToPath)
}

// SafeUnzip 把文件解压到给定的路径，附带安全检查
func (x *File) SafeUnzip(baseDirectory string) (err error) {
	saveToPath := filepath.Clean(filepath.Join(baseDirectory, x.Name))
	baseDirectory = filepath.Clean(baseDirectory)
	if IsZipSlip(baseDirectory, x.Name) {
		return BuildZipSlipError(x.File)
	}
	return x.Save(saveToPath)
}
