package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"

	legacy "awshowto/aws"
	. "awshowto/internal"
)

func main() {
	var isListRoles, isSts, isCreateRole, isDeleteRole, isAttachInlinePolicies bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&isListRoles, "list-roles", false, "AWS IAM list all roles.")
	flags.BoolVar(&isCreateRole, "create-role", false, "AWS IAM create role.")
	flags.BoolVar(&isDeleteRole, "delete-role", false, "AWS IAM delete role.")
	flags.BoolVar(&isAttachInlinePolicies, "attach-inline", false, "AWS IAM attach inline policies to a role.")
	flags.BoolVar(&isSts, "show-sts", false, "Show STS config.")

	_ = flags.Parse(os.Args[0:])
	args := flags.Args()

	if len(args) == 0 {
		printHelp(flags)
		return
	}

	cmp := legacy.New()

	if isListRoles {
		cmp.ListRoles()
	}

	if isCreateRole {
		cmp.CreateRole()
	}

	if isDeleteRole {
		cmp.DeleteRole()
	}

	if isAttachInlinePolicies {
		cmp.AttachInlinePolicies()
	}

	if isSts {
		acc, err := cmp.AccountAnalyzer()
		CheckError("sts", err)
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
