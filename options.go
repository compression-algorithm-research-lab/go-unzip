package unzip

// DefaultUnzipWorkerNum 如果没有指定的话，默认情况下解压使用的并发数是多少
const DefaultUnzipWorkerNum = 1

// Options 解压压缩文件时可以指定的各种选项
type Options struct {

	// 源压缩文件，必须是zip格式
	SourceZipFile string

	// 解压到的目标文件夹，必须是一个目录，如果不存在的话会自动创建，如果已经存在的话尽量为空，否则可能会被重复覆盖写文件
	DestinationDirectory string

	// 解压的时候使用的并发数
	WorkerNum int
}

func NewOptions() *Options {
	return &Options{
		WorkerNum: DefaultUnzipWorkerNum,
	}
}

// SetSourceZipFile 设置要解压的zip文件的路径
func (x *Options) SetSourceZipFile(sourceZipFile string) *Options {
	x.SourceZipFile = sourceZipFile
	return x
}

// SetDestinationDirectory 设置解压后输出的文件夹的路径
func (x *Options) SetDestinationDirectory(destinationDirectory string) *Options {
	x.DestinationDirectory = destinationDirectory
	return x
}

// SetWorkerNum 设置解压时使用到的并发数
func (x *Options) SetWorkerNum(workerNum int) *Options {
	x.WorkerNum = workerNum
	return x
}
