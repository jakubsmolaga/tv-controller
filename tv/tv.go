package tv

import (
	"log"

	"github.com/tarm/serial"
)

type TV struct {
	c serial.Config
	s serial.Port
}

func Init() {
	c := &serial.Config{Name: "COM45", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}
