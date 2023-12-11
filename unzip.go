package unzip

import (
	"archive/zip"
	"fmt"
	"github.com/golang-infrastructure/go-pointer"
	"sync"
)

// Unzip 用于封装解压缩的逻辑
type Unzip struct {
	options *Options
}

// New 从选项创建一个解压缩组件
func New(options *Options) *Unzip {
	return &Unzip{
		options: options,
	}
}

// FileHandler 用来处理解压出来的文件
type FileHandler func(file *File, options *Options) error

// SafeTraversal 安全的遍历压缩文件，遇到非法的或者错误的压缩文件时会自动检测报错
func (x *Unzip) SafeTraversal(handler FileHandler) (err error) {
	return x.Traversal(func(file *File, options *Options) error {
		if IsZipSlip(x.options.DestinationDirectory, file.Name) {
			return fmt.Errorf("zip slip, deny")
		}
		return handler(file, x.options)
	})
}

// Traversal 遍历zip文件
func (x *Unzip) Traversal(handler FileHandler) (err error) {

	// 参数检查
	if x.options.SourceZipFile == "" {
		return ErrSourceZipFileEmpty
	} else if x.options.WorkerNum == nil || pointer.FromPointer(x.options.WorkerNum) <= 0 {
		return ErrWorkerNumInvalid
	}

	// 打开压缩文件
	var r *zip.ReadCloser
	r, err = zip.OpenReader(x.options.SourceZipFile)
	if err != nil {
		return err
	}

	defer func() {
		// 如果有其它错误的话，优先返回其它错误的类型
		localError := r.Close()
		if err == nil {
			err = localError
		}
	}()

	// 初始化文件队列
	fileChannel := x.makeZipFileChannel(r.File)

	// 并发处理压缩文件中的每个文件
	var wg sync.WaitGroup
	for i := 0; i < pointer.FromPointer(x.options.WorkerNum); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChannel {
				// TODO 错误收集
				_ = handler(file, x.options)
			}
		}()
	}
	wg.Wait()
	return nil
}

// makeZipFileChannel 把zip数组转换为chan队列
func (x *Unzip) makeZipFileChannel(files []*zip.File) chan *File {
	fileChannel := make(chan *File, len(files))
	for _, f := range files {
		fileChannel <- &File{
			File: f,
		}
	}
	close(fileChannel)
	return fileChannel
}

//// CheckZipSlip 检查zip文件是否有zip slip漏洞
//func (x *Unzip) CheckZipSlip() []*zip.File {
//
//	x.Traversal(func(file *zip.File, options *Options) error {
//		if x.IsZipSlip(file) {
//
//		}
//	})
//}

// Unzip 解压文件到给定的目录
func (x *Unzip) Unzip() error {

	// 参数检查
	if x.options.WorkerNum == nil || pointer.FromPointer(x.options.WorkerNum) <= 0 {
		return ErrWorkerNumInvalid
	} else if x.options.SourceZipFile == "" {
		return ErrSourceZipFileEmpty
	} else if x.options.DestinationDirectory == "" {
		return ErrDestinationDirectoryEmpty
	}

	// 遍历压缩文件中的每个Entry，依此保存到磁盘上
	return x.SafeTraversal(func(file *File, options *Options) error {
		return file.Unzip(x.options.DestinationDirectory)
	})
}

// Files 读取压缩包中的所有文件和内容并返回，注意这个方法会实际解压缩文件，确保你机器的资源是足够的
func (x *Unzip) Files() (fileSlice []*File, err error) {

	var r *zip.ReadCloser
	r, err = zip.OpenReader(x.options.SourceZipFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		// 如果有其它错误的话，优先返回其它错误的类型
		localError := r.Close()
		if localError != nil && err == nil {
			err = localError
		}
	}()

	fileSlice = make([]*File, len(r.File))
	for i, f := range r.File {
		file := File{File: f}
		_, err := file.ReadBytes()
		if err != nil {
			return nil, err
		}
		fileSlice[i] = &file
	}

	return fileSlice, nil
}
