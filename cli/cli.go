package cli

import (
	"fmt"

	"github.com/sig00rd/simpleserver/file"
	"github.com/sig00rd/simpleserver/server"
)

const VERSION_COMMAND = "version"
const HELP_COMMAND = "help"
const RUN_COMMAND = "run"

const NO_SUBCOMMAND_MESSAGE = "Expected a subcommand"
const INVALID_SUBCOMMAND_MESSAGE = "Expected 'run', 'version' or 'help' subcommand"
const INVALID_RUN_COMMAND_MESSAGE = "Expected both --file flag and a file path, refer to help or readme"

const VERSION = "v1.0.0"

type CLI struct {
	Server server.Server
}

func (c *CLI) Parse(args []string) error {
	if len(args) <= 1 {
		return fmt.Errorf(NO_SUBCOMMAND_MESSAGE)
	}

	switch args[1] {
	case VERSION_COMMAND:
		version()

	case HELP_COMMAND:
		help(args[0])

	case RUN_COMMAND:
		if len(args) < 4 || args[2] != "--file" {
			return fmt.Errorf(INVALID_RUN_COMMAND_MESSAGE)
		}

		err := c.run(args[3])
		if err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf(INVALID_SUBCOMMAND_MESSAGE)
	}

	return fmt.Errorf("Encountered an error while parsing user input")
}

func (c *CLI) run(filePath string) error {
	err := file.InspectFile(filePath)
	if err != nil {
		return fmt.Errorf("Encountered an error while trying to fetch the file:\n %s", err)
	}

	return c.Server.Serve(filePath)
}

func version() {
	fmt.Println(VERSION)
}

func help(app string) {
	message := fmt.Sprintf(
		"Simple CLI tool for running a server serving html file\n"+
			"Usage:\n"+
			"  %s version to display app version\n"+
			"  %s help to display help\n"+
			"  %s run --file <path> to start HTTP server", app, app, app)
	fmt.Println(message)
}
