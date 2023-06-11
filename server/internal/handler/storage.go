package handler

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func HandleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4)
	_, _ = conn.Read(buf)

	cmdType := string(buf)
	if cmdType == "SENT" {
		receiveFile(ctx, conn)
	} else if cmdType == "RECV" {
		sendFile(ctx, conn)
	}

}

func receiveFile(ctx context.Context, conn net.Conn) {
	// read file id
	buf := make([]byte, 36)
	_, _ = conn.Read(buf)
	fileID := string(buf)

	// read file size
	buf = make([]byte, 8)
	_, _ = conn.Read(buf)

	var limit int64
	bf := bytes.NewBuffer(buf)
	_ = binary.Read(bf, binary.LittleEndian, &limit)

	path := fmt.Sprintf("./buckets/%s", fileID)
	file, err := os.Create(path)
	if err != nil {
		log.Println("error opening file:", err)
		return
	}
	defer file.Close()

	n, err := io.CopyN(file, conn, limit)
	if err != nil {
		log.Println("error receiving file:", err)
		return
	}

	log.Println("file received successfully", n)
}

func sendFile(ctx context.Context, conn net.Conn) {
	// read file id
	buf := make([]byte, 36)
	_, _ = conn.Read(buf)
	fileID := string(buf)

	path := fmt.Sprintf("./buckets/%s", fileID)
	file, err := os.Open(path)
	if err != nil {
		log.Println("error opening file:", err)
		return
	}
	defer file.Close()

	n, err := io.Copy(conn, file)
	if err != nil {
		log.Println("error sending file:", err)
	}

	log.Println("file sent successfully", n)
}
