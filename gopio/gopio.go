// Package gopio provides a struct and a few functions to interact with the GPIO pins on the Aria G25.
package gopio

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
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

// Directions for reading or writing.
const (
	In  = "in"
	Out = "out"
)

// Path to control interfaces. Not a constant so we can change this path in tests.
var ControlPath string = "/sys/class/gpio/"

// Return path to GPIO signal directory, like /sys/class/gpio/gpio114/
var getSignalPath = func(pin *Pin) string {
	return filepath.Join(ControlPath, strconv.Itoa(pin.KernelId), "/")
}

// A Pin represents a single Pin off the Aria G25.
type Pin struct {
	KernelId int // Kernel id of pin.
}

// Export control of GPIO to userspace.
func (pin *Pin) Export() {
	b := []byte(strconv.Itoa(pin.KernelId))
	err := ioutil.WriteFile(filepath.Join(ControlPath, "export"), b, 0755)

	if err != nil {
		log.Fatalf("Could not export pin with %b.", pin.KernelId)
	}
}

// Set direction of Pin to either "in" or "out" for reading or writing.
func (pin *Pin) setDirection(direction string) {
	path := filepath.Join(getSignalPath(pin), "direction")
	err := ioutil.WriteFile(path, []byte(direction), 0755)

	if err != nil {
		log.Fatalf("Could not open file for writing direction of %d: %s", pin.KernelId, err)
	}
}

// Return current value of Pin.
func (pin *Pin) Read() int {
	pin.setDirection(In)

	path := filepath.Join(getSignalPath(pin), "value")
	b, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Couldn't read GPIO: %v.", err)
	}

	if len(b) == 0 {
		log.Fatalf("Couldn't read state of pin from %v: file contains 0 bytes.", path)
	}

	value, _ := strconv.Atoi(string(b[0]))

	return value
}

// Write value to Pin.
func (pin *Pin) Write(value int) error {
	pin.setDirection(Out)
	path := filepath.Join(getSignalPath(pin), "value")

	b := []byte(strconv.Itoa(value))
	return ioutil.WriteFile(path, b, 0755)
}
