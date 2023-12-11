# go-unzip 

# 一、这是什么？优势是什么？

Go中用来解压zip文件的API，优势：

- 在原生的zip.File的基础上封装了易用API，一键解压
- 支持并发解压zip文件

# 二、安装

```bash
go get -u github.com/compression-algorithm-research-lab/go-unzip
```

# 三、API示例 

```go
package main

import "github.com/compression-algorithm-research-lab/go-unzip"

func main() {

	options := unzip.NewOptions().
		SetSourceZipFile("test_data/foo.zip").
		SetDestinationDirectory("test_data/foo").
		SetWorkerNum(100)
	err := unzip.New(options).Unzip()
	if err != nil {
		panic(err)
	}

}
```





