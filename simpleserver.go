package main

import (
	"log"
	"os"

	"github.com/sig00rd/simpleserver/cli"
	"github.com/sig00rd/simpleserver/server"
)

func main() {
	cli := &cli.CLI{&server.NormalServer{}}
	err := cli.Parse(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
