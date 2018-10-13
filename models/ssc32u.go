package models

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
	"log"
)

type SSC32U struct{
	port io.ReadWriteCloser
}

func InitSSC32U(dev string, baud uint) (*SSC32U, error) {
	options := serial.OpenOptions{
		PortName: dev,
		BaudRate: baud,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 8,
	}

	port, err := serial.Open(options)

	if err != nil {
		return nil, err
	}

	defer port.Close()

	ssc := &SSC32U{port: port}

	ssc.Write("VER\r")
	version, err := ssc.Read(128)

	if err != nil {
		return nil, err
	}

	log.Printf("Connected: %s", version)

	return ssc, nil
}

func (s SSC32U) Write(data string) (*int, error) {
	bytes := []byte(data)
	n, err := s.port.Write(bytes)

	if err != nil {
		return nil, err
	}

	return &n, nil
}

func (s SSC32U) Read(size uint) (string, error) {
	buffer := make([]byte, size)
	n, err := s.port.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}