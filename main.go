package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/dariusbakunas/go-bot-api/models"
	"github.com/dariusbakunas/go-bot-api/server"
	"log"
	"os"
)

func main() {
	parser := argparse.NewParser("go-bot-api", "Go Bot API to control SSC-32U")
	serialPort := parser.String("s", "serial-port", &argparse.Options{Required: false, Help: "USB Serial Port", Default: "/dev/ttyUSB0"})
	listenPort := parser.Int("p", "server-port", &argparse.Options{Required: false, Help: "Listening port", Default: 8080})
	baudRate := parser.Int("b", "baud-rate",  &argparse.Options{Required: false, Help: "Baud rate", Default: 115200})
	simulate := parser.Flag("d", "simulate", &argparse.Options{Required: false, Help: "Use fake serial port if set", Default: false})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	} else {
		ssc32u, err := models.InitSSC32U(*serialPort, uint(*baudRate), *simulate)
		defer ssc32u.Close()

		if err != nil {
			log.Fatalf("failed to open serial: %v", err)
		}

		server.Init(*listenPort, ssc32u)
	}
}