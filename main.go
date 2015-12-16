package clitesting

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

type Command struct {
	cmd       *exec.Cmd
	exitError *exec.ExitError
	executed  bool
	stdout    string
	stderr    string
}

var UninitializedCommand = errors.New("You need to run this command first")

func NewCommand(name string, arg ...string) *Command {
	return &Command{
		cmd: exec.Command(name, arg...),
	}
}

func (c *Command) Run() {
	var outBuf bytes.Buffer
	c.cmd = exec.Command("ls", "-l1")
	c.cmd.Stdout = &outBuf

	var errBuf bytes.Buffer
	c.cmd.Stderr = &errBuf

	if err := c.cmd.Run(); err != nil {
		log.Fatalf("ERR: %s", err)
		c.exitError = err.(*exec.ExitError)
	}
	c.stdout = string(outBuf.Bytes())
	c.stderr = string(errBuf.Bytes())
	c.executed = true
}

func (c *Command) Stdout() string {
	if !c.executed {
		log.Fatal(UninitializedCommand)
	}
	return c.stdout
}

func (c *Command) AssertSuccess(msg string) {
	if !c.executed {
		log.Fatal("Command needs to be executed before asserting anything")
	}
	if c.exitError != nil {
		log.Fatal(msg)
	}
}
