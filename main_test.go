package main

import "testing"

func TestSayHello(t *testing.T) {
	r := sayHello()
	if r != "Hello!" {
		t.Error("Not equal")
	}
}
