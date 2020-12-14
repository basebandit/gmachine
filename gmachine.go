// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

//G-Machine instruction-set with their respective OPCodes.
const (
	//OpHALT stops current execution.
	OpHALT = iota
	//OpNOOP does nothing.
	OpNOOP
	//OpINCA adds one to.
	OpINCA
	//OpDECA subtracts one from.
	OpDECA
	//OpSETA operates on an operand.
	OpSETA
)

//GMachine is a state-machine representing a virtual CPU.
type GMachine struct {
	//P is a 64 bit register. Used to store the memory address of the next instruction to execute.
	P uint64
	//A is a 64 bit register. Used to store results of computations. e.g. addition,subtraction.
	A uint64
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
		case OpINCA:
			//We add one to the value held by A,then increment the program counter.(PC)
			g.A++
			g.P++
		case OpDECA:
			// g.A = 2 //First set the register A to the value 2
			g.A-- //Then subtract one from the value of A register
			g.P++ //Then increment Program Counter
		case OpSETA:
			//trigger a memory fetch
			addr := g.P
			i := g.Memory[addr] //read contents of memory
			g.A = i             //put the read value into the A register
			g.P++               //increment Program Counter
		}
	}
}

//Sub subtracts 2 from 3. Sub first increments A three times setting it to value 3. Then decrements A two times resulting to the difference of subtracting 2 from 3.
func (g *GMachine) Sub() {
	g.Memory[0] = OpINCA // A = 0+ 1 = 1
	g.Memory[1] = OpINCA // A = 1 + 1 = 2
	g.Memory[2] = OpINCA // A = 2 + 1 = 3
	g.Memory[3] = OpDECA // A = 3 - 1 = 2
	g.Memory[4] = OpDECA // A = 2 - 1 = 1
	g.Run()
}
