package ftphandler

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type FTPHandler struct {
	rootDir string
}

func NewFTPHandler() *FTPHandler {
	return &FTPHandler{
		rootDir: "./ftp_root",
	}
}

func (h *FTPHandler) HandleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		message = strings.TrimSpace(message)

		switch {
		case strings.HasPrefix(message, "USER"):
			h.handleUser(message, writer)
		case strings.HasPrefix(message, "PASS"):
			h.handlePass(message, writer)
		case strings.HasPrefix(message, "CWD"):
			h.handleCwd(message, writer)
		case strings.HasPrefix(message, "LIST"):
			h.handleList(message, writer)
		}
	}
}


func (h *FTPHandler) handleUser(message string, writer *bufio.Writer){
	writer.WriteString("331 User name okay, need password\r\n")
	writer.Flush()
}

func (h *FTPHandler) handlePass(message string, writer *bufio.Writer){
	writer.WriteString("230 User Logged in \r\n")
	writer.Flush()
}

func (h *FTPHandler) handleCwd(message string, writer *bufio.Writer){
	dir := strings.TrimSpace(message[4:])
	path := filepath.Join(h.rootDir, dir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		writer.WriteString("550 Directory not found \r\n")
		writer.Flush()
		return
	}

	writer.WriteString("250 Directory changed\r\n")
	writer.Flush()
}

func (h *FTPHandler) handleList(message string, writer *bufio.Writer){
	path := filepath.Join(h.rootDir, ".")

	files, err := os.ReadDir(path)
	if err != nil {
		writer.WriteString("550 Directory not found\r\n")
		writer.Flush()
		return
	}

	writer.WriteString("150 Here comes the directory listing\r\n")
	for _, file := range files {
		writer.WriteString(fmt.Sprintf("%s\r\n", file.Name()))
	}
	writer.WriteString("226 Directory listing ended\r\n")
	writer.Flush()
}

func (h *FTPHandler) handleRetr(message string, writer *bufio.Writer){
	filePath := strings.TrimSpace(message[5:])
	path := filepath.Join(h.rootDir, filePath)

	file, err := os.Open(path)
	if err != nil {
		writer.WriteString("550 File not found \r\n")
		writer.Flush()
		return
	}
	defer file.Close()

	writer.WriteString("150 Opening BINARY mode data connection\r\n")
	writer.Flush()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		writer.Write(buf[:n])
		writer.Flush()
	}

	writer.WriteString("226 Transfer complete\r\n")
	writer.Flush()
}

func (h *FTPHandler) handleStor(message string, reader *bufio.Reader, writer *bufio.Writer) {
	filePath := strings.TrimSpace(message[5:])
	path := filepath.Join(h.rootDir, filePath)

	file, err := os.Create(path)
	if err != nil {
		writer.WriteString("550 File cannot be created\r\n")
		writer.Flush()
		return
	}
	defer file.Close()

	writer.WriteString("150 Opening BINARY mode data connection\r\n")
	writer.Flush()

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}
		file.Write(buf[:n])
	}

	writer.WriteString("226 Transfer complete\r\n")
	writer.Flush()
}