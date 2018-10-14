package models

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"log"
)

type SSC32U struct{
	port io.ReadWriteCloser
	simulate bool
}

func InitSSC32U(dev string, baud uint, simulate bool) (*SSC32U, error) {
	options := serial.OpenOptions{
		PortName: dev,
		BaudRate: baud,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 8,
	}

	var port io.ReadWriteCloser
	var err error

	if !simulate {
		port, err = serial.Open(options)

		if err != nil {
			return nil, err
		}

		defer port.Close()
	}

	ssc := &SSC32U{port: port, simulate: simulate}

	ssc.Write("VER\r")
	version, err := ssc.Read(128)

	if err != nil {
		return nil, err
	}

	log.Printf("Connected: %s", version)

	return ssc, nil
}

func (s SSC32U) Write(data string) (int, error) {
	if s.simulate {
		message := fmt.Sprintf("Writing: %s", data)
		log.Print(message)
		return len(message), nil
	}

	bytes := []byte(data)
	n, err := s.port.Write(bytes)

	if err != nil {
		return -1, err
	}

	return n, nil
}

func (s SSC32U) Read(size uint) (string, error) {
	if s.simulate {
		return "NOP", nil
	}

	buffer := make([]byte, size)
	n, err := s.port.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}