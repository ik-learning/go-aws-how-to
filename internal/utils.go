package internal

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var supported = map[string]bool{
	"ec2":         true,
	"lambda":      true,
	"rds":         true,
	"redshift":    true,
	"kafka":       true,
	"opensearch":  true,
	"elasticache": true,
	"emr":         true,
}

func MatchSupportedTypes(resources []string) error {
	for _, resource := range resources {
		if _, found := supported[resource]; !found {
			msg := fmt.Sprintf("Resource %v not supported", resource)
			return errors.New(msg)
		}
	}
	return nil
}

func CheckError(message string, err error) {
	if err != nil {
		ExitErrorf(message, err)
	}
}

func ExitErrorf(msg string, args ...interface{}) {
	if args == nil {
		_, _ = fmt.Fprint(os.Stderr, color.RedString("\t❌ %s\n", msg))
	} else {
		_, _ = fmt.Fprint(os.Stderr, color.RedString("\t❌ %s\n\t %s\n", msg, args))
	}
	os.Exit(1)
}

func OutputColorizedMessage(clr string, message string) {
	switch clr {
	case "blue":
		_, _ = fmt.Fprint(os.Stdout, color.BlueString(fmt.Sprintf("\t%s\n", message)))
	case "green":
		_, _ = fmt.Fprint(os.Stdout, color.GreenString(fmt.Sprintf("\t%s\n", message)))
  case "yellow":
		_, _ = fmt.Fprint(os.Stdout, color.YellowString(fmt.Sprintf("\t%s\n", message)))
	default:
		_, _ = fmt.Fprint(os.Stdout, color.HiRedString(fmt.Sprintf("\t%s\n", message)))
	}
}
