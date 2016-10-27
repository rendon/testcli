package testcli

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSetEnv(t *testing.T) {
	c := Command("/bin/sh", "-c", "echo -n $FOO")
	c.SetEnv([]string{"FOO=bar"})
	c.Run()
	if c.Failure() {
		t.Fatalf("Expected to succeed, but failed")
	}

	if c.stdout != "bar" {
		t.Log(c.stdout)
		t.Fatal("stdout failed to include input")
	}
}

func TestSetStdin(t *testing.T) {
	buf := bytes.NewBufferString("foo\n")
	c := Command("cat")
	c.SetStdin(buf)
	c.Run()
	if c.Failure() {
		t.Fatalf("Expected to succeed, but failed")
	}

	if c.stdout != "foo\n" {
		t.Fatal("stdout failed to include input")
	}
}

func TestFailedRun(t *testing.T) {
	c := Command("myunknowncommand")
	c.Run()
	if !c.Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
}

func TestPackageFailedRun(t *testing.T) {
	Run("myunknowncommand")
	if !Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
}

func TestSuccessfulRun(t *testing.T) {
	c := Command("whoami")
	c.Run()
	if !c.Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
	}
}

func TestPackageSuccessfulRun(t *testing.T) {
	Run("whoami")
	if !Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", Error())
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

func TestPackageStdout(t *testing.T) {
	user := os.Getenv("USER")
	Run("whoami")
	if !StdoutContains(user) {
		t.Fatalf("Expected %q to contains %q", Stdout(), user)
	}

	// testing case insensitiveness
	user = strings.ToUpper(user)
	if !StdoutContains(user) {
		t.Fatalf("Expected %q to contains %q", Stdout(), user)
	}
}

func TestStderr(t *testing.T) {
	c := Command("cp")
	c.Run()
	if !c.Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
	if c.Stderr() == "" {
		t.Fatalf("Expected %q NOT to be empty", c.Stderr())
	}

	// testing case insensitiveness
	if !c.StderrContains("MISSING") {
		t.Fatalf("Expected %q to contains %q", c.Stderr(), "MISSING")
	}
}

func TestPackageStderr(t *testing.T) {
	Run("cp")
	if !Failure() {
		t.Fatalf("Expected to fail, but succeeded")
	}
	if !StderrContains("missing") {
		t.Fatalf("Expected %q to contains %q", Stderr(), "missing")
	}

	// testing case insensitiveness
	if !StderrContains("MISSING") {
		t.Fatalf("Expected %q to contains %q", Stderr(), "MISSING")
	}
}

func TestStdoutMatches(t *testing.T) {
	regex := "(/[^/]*)+"
	c := Command("whoami")
	c.Run()
	if c.StdoutMatches(regex) {
		t.Fatalf("Expected %q NOT to match %q", c.Stdout(), regex)
	}

	c = Command("pwd")
	c.Run()
	if !c.StdoutMatches(regex) {
		t.Fatalf("Expected %q to match %q", c.Stdout(), regex)
	}
}

func TestPackageStdoutMatches(t *testing.T) {
	regex := "(/[^/]*)+"
	Run("whoami")
	if StdoutMatches(regex) {
		t.Fatalf("Expected %q NOT to match %q", Stdout(), regex)
	}

	Run("pwd")
	if !StdoutMatches(regex) {
		t.Fatalf("Expected %q to match %q", Stdout(), regex)
	}
}

func TestStderrMatches(t *testing.T) {
	regex := "(/[^/]*)+"
	c := Command("cp")
	c.Run()
	if c.StderrMatches(regex) {
		t.Fatalf("Expected %q NOT to match %q", c.Stderr(), regex)
	}
	regex = "cp.*operand"
	if !c.StderrMatches(regex) {
		t.Fatalf("Expected %q to match %q", c.Stderr(), regex)
	}
}

func TestPackageStderrMatches(t *testing.T) {
	regex := "(/[^/]*)+"
	Run("cp")
	if StderrMatches(regex) {
		t.Fatalf("Expected %q NOT to match %q", Stderr(), regex)
	}
	regex = ".*cp.*"
	if !StderrMatches(regex) {
		t.Fatalf("Expected %q to match %q", Stderr(), regex)
	}
}
