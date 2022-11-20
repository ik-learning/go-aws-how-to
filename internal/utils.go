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
	_, _ = fmt.Fprint(os.Stderr, color.RedString("❌  %s: %s\n", msg, args))
	os.Exit(1)
}

func OutputColorizedMessage(clr string, message string) {
	switch clr {
	case "blue":
		_, _ = fmt.Fprint(os.Stdout, color.BlueString(message))
	case "green":
		_, _ = fmt.Fprint(os.Stdout, color.GreenString(message))
	default:
		_, _ = fmt.Fprint(os.Stdout, color.HiRedString(message))
	}
}
