package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// File 用于表示一个zip压缩文件，在原生的zip file上封装了一些基础操作
type File struct {

	// 底层基于的zip file
	*zip.File

	// 文件的字节内容，对读文件内容解压缩做一个缓存
	fileUnzipCacheBytes []byte
}

// NewFile 从压缩文件创建文件
func NewFile(zipFile *zip.File) *File {
	return &File{
		File: zipFile,
	}
}

// ReadBytes 读取此压缩文件的字节数组
func (x *File) ReadBytes() (bytes []byte, returnError error) {

	// 如果之前已经读取过了，则返回缓存，不再重复读取
	if x.fileUnzipCacheBytes != nil {
		return x.fileUnzipCacheBytes, nil
	}

	open, returnError := x.Open()
	if returnError != nil {
		return nil, returnError
	}
	defer func() {
		err := open.Close()
		// 避免覆盖掉更重要的错误
		if err != nil && returnError == nil {
			returnError = err
		}
	}()
	x.fileUnzipCacheBytes, returnError = io.ReadAll(open)
	return x.fileUnzipCacheBytes, returnError
}

// Save 把此文件解压并保存到硬盘上给定的位置
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
