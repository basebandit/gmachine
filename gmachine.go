// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

//GMachine is a state-machine representing a virtual CPU.
type GMachine struct {
	//P is a 64 bit register. Used to store the memory address of the next instruction to execute.
	P uint64
	//Memory is an array of 64 bit words that represent the memory of a G-Machine.
	Memory []uint64
}

//New initializes a GMachine instance to its default state. Returns a pointer to a G-Machine instance.
func New() *GMachine {
	gm := GMachine{
		Memory: make([]uint64, DefaultMemSize),
	}

	return &gm
}
