package main

import "github.com/compression-algorithm-research-lab/go-unzip"

func main() {

	options := unzip.NewOptions().SetSourceZipFile("test_data/foo.zip").SetDestinationDirectory("test_data/foo")
	err := unzip.New(options).Unzip()
	if err != nil {
		panic(err)
	}

}
