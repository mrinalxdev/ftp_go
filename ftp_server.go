package main

import (
	"fmt"
	"log"

	ftp_handler "github.com/mrinalxdev/ftp_go/ftp_handler"
	ftp_listener "github.com/mrinalxdev/ftp_go/ftp_listener"
)

const (
	ADDRESS = "localhost:2121"
)

func main() {
	listener, err := ftp_listener.NewFTPListener(ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	handler := ftp_handler.NewFTPHandler()

	fmt.Println("FTP Server listening on", ADDRESS)
	if err := listener.Listen(handler); err != nil {
		log.Fatal(err)
	}
}
