package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

// hashworkTests known values
var hashworkTests = []struct {
	message  string
	rounds   int
	expected string
}{
	{"abc", 1, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
	{"abc", 10, "10e286f907c0fe9f02cea3864cbaec04ae47e2c0a13b60473bc9968a4851b219"},
}

// TesthashworkOutPut
func TestHashworkOutPut(t *testing.T) {
	message := []byte("abc")
	hash := hashwork(message, 1)
	o := sha256.Sum256(message)
	expected := o[:]
	if string(hash) != string(expected) {
		t.Errorf("hashwork: hash failed %v", message)
	}
}

func TestHashwork(t *testing.T) {
	for _, mt := range hashworkTests {
		m := []byte(mt.message)
		r := hashwork(m, mt.rounds)
		rs := fmt.Sprintf("%x", r)
		if rs != mt.expected {
			t.Errorf("%s %s", mt.expected, rs)
		}
	}
}
