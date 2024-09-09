package main

import "testing"

func TestSayHello(t *testing.T) {
	r := sayHello()
	if r != "Hello world!" {
		t.Error("Not equal")
	}
}
