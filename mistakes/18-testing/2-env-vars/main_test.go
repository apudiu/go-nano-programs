package main

import (
	"os"
	"testing"
)

func main() {}

func TestMainFn(t *testing.T) {
	if os.Getenv("INTEGRATION") != "1" {
		t.Skip("skipping integration test")
	}
}
