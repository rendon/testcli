package clitesting

import (
	"testing"
)

func TestRun(t *testing.T) {
	c := NewCommand("ls", "-l1")
	c.Run()
	if c.Stdout() == "" {
		t.Fatal("Expected non-empty string")
	}
}
