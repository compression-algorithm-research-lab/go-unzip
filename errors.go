package unzip

import "errors"

var (

	// ErrWorkerNumInvalid 解压zip文件时的并发数必须大于0，否则没法搞了
	ErrWorkerNumInvalid = errors.New("options.WorkerNum it has to be greater than 0")

	// ErrSourceZipFileEmpty 源zip文件不能为空，否则会返回此错误
	ErrSourceZipFileEmpty = errors.New("options.SourceZipFile can not empty")

	// ErrDestinationDirectoryEmpty 解压到的目录不能为空
	ErrDestinationDirectoryEmpty = errors.New("options.DestinationDirectory can not empty")
)
