package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Compute struct {
	context.Context
	aws.Config
}

func ExternalNew(ctx context.Context, cfg aws.Config) *Compute {
	return &Compute{ctx, cfg}
}

func New() *Compute {
	ctx := context.TODO()
	var cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return &Compute{ctx, cfg}
}
