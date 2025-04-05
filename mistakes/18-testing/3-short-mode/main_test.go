package main

import (
    "testing"
)

func main() {}

func TestLongRunning(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping long-running test")
    }
}
