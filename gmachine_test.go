package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	g := gmachine.New()
	g.Run()
	var wantP uint64 = 1
	if wantP != g.P {
		t.Errorf("want P == %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpNOOP
	g.Run()

	var wantP uint64 = 2
	if wantP != g.P {
		t.Errorf("want P == %d, got %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpINCA
	g.Run()

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpINCA
	g.Memory[1] = gmachine.OpINCA
	g.Memory[2] = gmachine.OpDECA
	g.Run()

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

//TestCALC calculate the result of subtracting 2 from 3
func TestCALC(t *testing.T) {
	g := gmachine.New()
	g.Sub()

	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}
