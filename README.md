# clitesting
CLI testing package for the Go language.

Developing a command line application? Wanna be able to test your app from the outside? If the answer is Yes to at least one of the questions, keep reading.

When using Ruby I use [aruba](https://github.com/cucumber/aruba) for testing command line applications, in Go I still can use aruba, but it"s awkward to bring Ruby and it's artillery only to test my app.

`clitesting` is a wrapper around [os.exec](https://golang.org/pkg/os/exec/) to test CLI apps in Go lang, minimalistic, so you can do your tests with [testing](https://golang.org/pkg/testing/) or any other testing framework.


## Greetings app
main\_test.go
```go
package main

import (
    "testing"

    "github.com/rendon/clitesting"
)

func TestGreetings(t *testing.T) {
    // make sure to execute `go install` before
    c := clitesting.NewCommand("greetings")
    c.Run()
    if !c.Success() {
        t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
    }

    if !c.StdoutContains("Hello?") {
        t.Fatalf("Expected %q to contain %q", c.Stdout(), "Hello?")
    }
}

func TestGreetingsWithName(t *testing.T) {
    // make sure to execute `go install` before
    c := clitesting.NewCommand("greetings", "--name", "John")
    c.Run()
    if !c.Success() {
        t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
    }

    if !c.StdoutContains("Hello John!") {
        t.Fatalf("Expected %q to contain %q", c.Stdout(), "Hello John!")
    }
}
```


main.go
```go
package main

import (
    "fmt"
    "os"

    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "cli"
    app.Usage = "CLI app"
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "name",
            Usage: "User name",
        },
    }
    app.Action = func(c *cli.Context) {
        if c.String("name") != "" {
            fmt.Printf("Hello %s!\n", c.String("name"))
        } else {
            fmt.Printf("Hello? Anyone?\n")
        }
    }

    app.Run(os.Args)
}
```
