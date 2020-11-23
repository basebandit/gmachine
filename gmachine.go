// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

//G-Machine instruction-set with their respective OPCodes.
const (
	//OpHALT stops current execution.
	OpHALT = 0
	//OpNOOP does nothing.
	OpNOOP = 1
)

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

//Run executes an instruction on the new G-Machine instance.
func (g *GMachine) Run() {
	for { //Fetch-Execute Cycle
		//Step 1: Fetch the next instruction from Memory. Look at the P register to see what memory address it
		// contains, and read the instruction at that address.
		addr := g.P
		i := g.Memory[addr]

		switch i {
		case OpHALT: //Step 3: Increment the P register (PC) so that it points to the next memory address to read from.
			g.P++
			return
		case OpNOOP:
			//We do nothing and increment the program counter (PC).
			g.P++
		}
	}
}
