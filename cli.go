package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// version is the current version of the CLI.
var version string = "0.1.0"

// Cli represents the command-line interface.
type Cli struct {
	OutStream, ErrStream io.Writer
}

// CliArgs holds the parsed command-line arguments.
type CliArgs struct {
	ShowHelp    bool
	ShowVersion bool
	SearchWords []string
}

// ParseArgs parses command-line arguments.
func ParseArgs(args []string) (CliArgs, error) {
	var cliArgs CliArgs

	flags := flag.NewFlagSet("gh-kw", flag.ContinueOnError)
	flags.BoolVar(&cliArgs.ShowHelp, "h", false, "Show help message")
	flags.BoolVar(&cliArgs.ShowHelp, "help", false, "Show help message")
	flags.BoolVar(&cliArgs.ShowVersion, "v", false, "Show version")
	flags.BoolVar(&cliArgs.ShowVersion, "version", false, "Show version")

	if err := flags.Parse(args); err != nil {
		// nolint: wrapcheck
		return cliArgs, err
	}

	cliArgs.SearchWords = flags.Args()

	return cliArgs, nil
}

// Run parses command-line arguments and executes the appropriate action.
func (cli *Cli) Run() int {
	cliArgs, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(cli.ErrStream, err)
		return 1
	}

	if cliArgs.ShowHelp {
		cli.usage()
		return 0
	}

	if cliArgs.ShowVersion {
		fmt.Fprintf(cli.OutStream, "gh-dot-tmpl version %s\n", version)
		return 0
	}

	if len(cliArgs.SearchWords) == 0 {
		fmt.Fprintf(cli.ErrStream, "Error: No search word provided\n")
		cli.usage()

		return 1
	}

	searchWords := cliArgs.SearchWords

	sr, err := Search(searchWords)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "Error: %s\n", err)
		return 1
	}

	// Sort search results by search count in descending order.
	sortedResults := SortSearchResults(sr)

	// Output search results in a table.
	Output(sortedResults, cli.OutStream)

	return 0
}

// usage prints the help message.
func (cli *Cli) usage() {
	fmt.Fprintf(cli.OutStream, `Usage: gh kw [options] [search_word...]

Options:
  -h, --help       Show help message
  -v, --version    Show version

Arguments:
  search_word...  Names of the search words
`)
}
