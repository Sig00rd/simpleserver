## Recruitment task from a certain company

Simple CLI tool to start a server serving static html file

## Run tests

In the project root folder:

`go test ./...`

or if you're using richgo:

`richgo test ./...`

for verbose output, add '-v' option:

`[rich]go test -v ./...`

## Build app

cd into project root folder, then:

` go build simpleserver.go`

## Run app

With app built:
`./simpleserver version` to display info about app version

`./simpleserver help` to display help

`./simpleserver run --file <path>` to start HTTP server serving html file found under specified path
