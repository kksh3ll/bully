package main

import (
	"testing"
	"net"
	//"runtime"
	"fmt"
	"time"
)

func TestSingleBully(t *testing.T) {
	//runtime.GOMAXPROCS(2)
	ln, err := net.Listen("tcp", ":8801")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	bully := NewBully(ln, nil)
	err = bully.AddCandidate("127.0.0.1:8801", nil, 3 * time.Second)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	candy := bully.CandidateList()
	if len(candy) != 1 {
		t.Errorf("Wrong!")
	}
	for _, c := range candy {
		fmt.Printf("%v; %v\n", c.Addr, c.Id)
	}
	ln.Close()
}

func TestDoubleBully(t *testing.T) {
	aliceLn, err := net.Listen("tcp", ":8802")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	alice := NewBully(aliceLn, nil)

	bobLn, err := net.Listen("tcp", ":8082")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if bobLn == nil {
		t.Errorf("WTF\n")
	}
	bobAddr := "127.0.0.1:8082"
	bob := NewBully(bobLn, nil)

	err = alice.AddCandidate(bobAddr, nil, 3 * time.Second)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	candy := bob.CandidateList()
	if len(candy) != 2 {
		t.Errorf("Should be 2 candidate!")
	}
	fmt.Printf("Candidates in bob's list:\n")
	for _, c := range candy {
		fmt.Printf("%v; %v\n", c.Addr, c.Id)
	}
	aliceLn.Close()
	bobLn.Close()
}

