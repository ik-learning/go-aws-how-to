package main

import (
  "fmt"
  "os"
  "strings"

  flag "github.com/spf13/pflag"
)

func main() {
  var isListRoles bool

  flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
  flags.Usage = func() {
    printHelp(flags)
  }

  flags.BoolVar(&isListRoles, "list-roles", false, "List all AWS IAM roles.")

  _ = flags.Parse(os.Args[0:])
	args := flags.Args()

	if len(args) == 0 {
		printHelp(flags)
		return
	}

  if isListRoles {
    print("INSIDE")
  }

}

func printHelp(fs *flag.FlagSet) {
  _, _ = fmt.Fprintf(os.Stderr, "\n"+strings.TrimSpace(help)+"\n")
  fs.PrintDefaults()
}

const help = `
awshowto - how to code AWS resources with aws-go-sdk.

USAGE:
    $ awshowto [flags]

FLAGS:
`
