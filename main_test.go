package testcli

import (
	"os"
	"strings"
	"testing"
)

func TestFailedRun(t *testing.T) {
	c := Command("myunknowncommand")
	c.Run()
	if !c.Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
}

func TestSuccessfulRun(t *testing.T) {
	c := Command("whoami")
	c.Run()
	if !c.Success() {
		t.Fatal("Expected to succeed, but failed with error: %s", c.Error())
	}
}

func TestStdout(t *testing.T) {
	user := os.Getenv("USER")
	c := Command("whoami")
	c.Run()
	if !c.StdoutContains(user) {
		t.Fatalf("Expected %q to contains %q", c.Stdout(), user)
	}

	// testing case insensitiveness
	user = strings.ToUpper(user)
	if !c.StdoutContains(user) {
		t.Fatalf("Expected %q to contains %q", c.Stdout(), user)
	}
}

func TestStderr(t *testing.T) {
	c := Command("cp")
	c.Run()
	if !c.Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
	if !c.StderrContains("missing") {
		t.Fatalf("Expected %q to contains %q", c.Stderr(), "missing")
	}

	// testing case insensitiveness
	if !c.StderrContains("MISSING") {
		t.Fatalf("Expected %q to contains %q", c.Stderr(), "MISSING")
	}
}
