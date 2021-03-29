package cli

import (
	"testing"

	"github.com/sig00rd/simpleserver/helpers"
)

func TestParse(t *testing.T) {

	t.Run("should return an error when no subcommand is provided", func(t *testing.T) {
		args := []string{"./simpleserver"}
		cli := CLI{}

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, NO_SUBCOMMAND_MESSAGE)
	})

	t.Run("should return an error when unknown subcommand is issued", func(t *testing.T) {
		args := []string{"./simpleserver", "somecommandthatisnotimplemented"}
		cli := CLI{}

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_SUBCOMMAND_MESSAGE)
	})

	t.Run("should return an error when run is used without providing a file", func(t *testing.T) {
		args := []string{"./simpleserver", "run", "--file"}
		cli := CLI{}

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_RUN_COMMAND_MESSAGE)
	})

	t.Run("should return an error when run is used without a --file flag", func(t *testing.T) {
		args := []string{"./simpleserver", "run", "../test_files/correct.html"}
		cli := CLI{}

		got := cli.Parse(args)
		helpers.AssertError(t, got)
		helpers.AssertErrorMessage(t, got, INVALID_RUN_COMMAND_MESSAGE)
	})
}
