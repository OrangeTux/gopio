package gopio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
)

// Kernel id's of all pins, see http://www.acmesystems.it/pinout_ariag25.
const (
	N2 = iota + 96
	N3
	N4
	N5
	N6
	N7
	N8
	N9
	N10
	N11
	N12
	N13
	N14
	N15
	N16
	N17
	N18
	N19
	N20
	N21
	N22
	N23
	E2
	E3
	E4
	E5
	E6
	E7
	E8
	E9
	E10
	E11
	S23 = iota + 35
	S22
	S21
	S20
	S19
	S18
	S17
	S16
	S15
	S14
	S13
	S12
	S11
	S10
	S9
	S8
	S7
	S6
	S5
	S4
	S3
	S2
	S1
	W9
	W10
	W11
	W12
	W13
	W14
	W15
	W16
	W17
	W18
	W19
	W20
	W21
	W22
	W23
)

// Values that can be read and written from/to Pin.
const (
	LOW  = 0
	HIGH = 1
)

// Directions for reading or writing.
const (
	IN  = "in"
	OUT = "out"
)

// A Pin represents a single Pin off the Aria G25.
type Pin struct {
	kernelID  uint64 // Kernel id of pin.
	direction string
}

func (pin *Pin) export() {
	b := make([]byte, 1)
	binary.PutUvarint(b, pin.kernelID)
	err := ioutil.WriteFile("/sys/class/gpio/export", b, 0644)

	if err != nil {
		log.Fatal("Could not export pin with %b.", pin.kernelID)
	}
}

func (pin *Pin) setDirection(direction string) {
	path := fmt.Sprintf("/sys/class/gpio/gpio%b/direction", pin.kernelID)
	err := ioutil.WriteFile(path, []byte(direction), 0644)

	if err != nil {
		log.Fatal("Could not open file for writing direction of %b.", pin.kernelID)
	}

	pin.direction = direction
}

// Return current value of Pin.
func (pin Pin) Read() (value int) {
	if pin.direction != IN {
		pin.setDirection(IN)
	}

	path := fmt.Sprintf("/sys/class/gpio/gpio%b", pin.kernelID)
	b, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Couldn't read GPIO.")
	}

	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.LittleEndian, &value)

	if err != nil {
		log.Fatal("Couldn't read value from GPIO.")
	}

	return value
}