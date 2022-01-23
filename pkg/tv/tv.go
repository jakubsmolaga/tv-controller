package tv

import (
	"log"
	"net/http"

	"github.com/tarm/serial"
)

type TV struct {
	c serial.Config
	s serial.Port
}

func Init(port string) TV {
	c := &serial.Config{Name: port, Baud: 9600, Parity: serial.ParityNone, Size: 8, StopBits: 1}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return TV{c: *c, s: *s}
}

func (tv TV) sendCommand(c1 byte, c2 byte, d byte) {
	_, err := tv.s.Write([]byte{c1, c2, ' ', 0x00, ' ', d, '\r', '\n'})
	if err != nil {
		log.Fatal(err)
	}
}

func (tv TV) TurnOff(w http.ResponseWriter, r *http.Request) {
	tv.sendCommand('k', 'a', 0x00)
	w.Write([]byte("Ok"))
}

func (tv TV) TurnOn(w http.ResponseWriter, r *http.Request) {
	tv.sendCommand('k', 'a', 0x01)
	w.Write([]byte("Ok"))
}

func (tv TV) Reboot(w http.ResponseWriter, r *http.Request) {
	tv.sendCommand('k', 'a', 0x02)
	w.Write([]byte("Ok"))
}

func (tv TV) SelectHDMI1(w http.ResponseWriter, r *http.Request) {
	tv.sendCommand('x', 'b', 0x90)
	w.Write([]byte("Ok"))
}

func (tv TV) SelectDisplayPortPC(w http.ResponseWriter, r *http.Request) {
	tv.sendCommand('x', 'b', 0xd0)
	w.Write([]byte("Ok"))
}
