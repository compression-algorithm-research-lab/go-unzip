package unzip

type Options struct {

	// 源压缩文件
	SourceZipFile string

	// 解压到的目标文件
	DestinationDirectory string

	// 并发数
	WorkerNum int
}

func NewOptions() *Options {
	return &Options{
		WorkerNum: 1,
	}
}

func (x *Options) SetSourceZipFile(sourceZipFile string) *Options {
	x.SourceZipFile = sourceZipFile
	return x
}

func (x *Options) SetDestinationDirectory(destinationDirectory string) *Options {
	x.DestinationDirectory = destinationDirectory
	return x
}

func (x *Options) SetWorkerNum(workerNum int) *Options {
	x.WorkerNum = workerNum
	return x
}
