package main

import (
	"github.com/ABAlosaimi/BroCachi/services/cacheReader"
)

func main() {
	grpcServer := cachereader.NewCacheReaderServer(":9000")
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}
}