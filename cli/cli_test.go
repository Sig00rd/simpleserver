package cli

import (
	"testing"

	"github.com/sig00rd/simpleserver/helpers"
)

type ServerStub struct{}

func (ss *ServerStub) Serve(filePath string) error {
	return nil
}

func TestParse(t *testing.T) {

	t.Run("should return an error when no subcommand is provided", func(t *testing.T) {
		args := []string{"./simpleserver"}
		cli := newCLI()

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, NO_SUBCOMMAND_MESSAGE)
	})

	t.Run("should return an error when unknown subcommand is issued", func(t *testing.T) {
		args := []string{"./simpleserver", "somecommandthatisnotimplemented"}
		cli := newCLI()

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_SUBCOMMAND_MESSAGE)
	})

	t.Run("should return an error when run is used without providing a file", func(t *testing.T) {
		args := []string{"./simpleserver", "run", "--file"}
		cli := newCLI()

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_RUN_COMMAND_MESSAGE)
	})

	t.Run("should return an error when run is used without a --file flag", func(t *testing.T) {
		args := []string{"./simpleserver", "run", "../test_files/valid.html"}
		cli := newCLI()

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_RUN_COMMAND_MESSAGE)
	})

	t.Run("should not return an error when run is used with both --file flag and correct file path", func(t *testing.T) {
		args := []string{"./simpleserver", "run", "--file", "../test_files/valid.html"}
		cli := newCLI()

		got := cli.Parse(args)
		helpers.AssertNoError(t, got)
	})
}

func newCLI() *CLI {
	return &CLI{&ServerStub{}}
}
