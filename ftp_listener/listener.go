package ftp_listener

import (
	"net"
	ftp_handler "github.com/mrinalxdev/ftp_go/ftp_handler"
)

type FTPListener struct {
	address string
	handler *ftp_handler.FTPHandler
	ln      net.Listener
}

func NewFTPListener(address string) (*FTPListener, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &FTPListener{
		address: address,
		ln:      ln,
	}, nil
}

func (l *FTPListener) Listen(handler *ftp_handler.FTPHandler) error {
	l.handler = handler

	for {
		conn, err := l.ln.Accept()
		if err != nil {
			return err
		}

		go l.handler.HandleConnection(conn)
	}
}
