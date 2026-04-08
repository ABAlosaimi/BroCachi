package cachereader


func main() {
	grpcServer := NewCacheReaderServer(":50051")
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}

}