package unzip

import "errors"

var (

	// ErrWorkerNumInvalid 并发数必须大于0，否则没法搞了
	ErrWorkerNumInvalid = errors.New("options.WorkerNum it has to be greater than 0")

	// ErrSourceZipFileEmpty 源zip文件不能为空
	ErrSourceZipFileEmpty = errors.New("options.SourceZipFile can not empty")

	// ErrDestinationDirectoryEmpty 解压到的目录不能为空
	ErrDestinationDirectoryEmpty = errors.New("options.DestinationDirectory can not empty")
)
