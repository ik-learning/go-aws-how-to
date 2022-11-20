package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	flag "github.com/spf13/pflag"

	legacy "awshowto/aws"
	. "awshowto/internal"
)

func main() {
	var isListRoles, isSts bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&isListRoles, "list-roles", false, "List all AWS IAM roles.")
	flags.BoolVar(&isSts, "show-sts", false, "Show STS config.")

	_ = flags.Parse(os.Args[0:])
	args := flags.Args()

	if len(args) == 0 {
		printHelp(flags)
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	CheckError("Error: failed to call", err)

	cmp := legacy.New(cfg)

	if isListRoles {
		print("INSIDE LIST ROLES")
	}

	if isSts {
		acc, err := cmp.AccountAnalyzer()
		CheckError("iam", err)
		OutputColorizedMessage("blue", fmt.Sprintf("\tAccount:: %s. Aliases: %s\n", acc.Account, acc.Aliases))
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
