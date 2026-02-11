package main

import "testing"

func TestGetenv(t *testing.T) {
	if getenv("NO_SUCH_KEY", "x") != "x" {
		t.Fatal("default not returned")
	}
}
