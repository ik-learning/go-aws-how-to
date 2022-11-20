package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Compute struct {
	context.Context
	aws.Config
}

func ExternalNew(ctx context.Context, cfg aws.Config) *Compute {
	return &Compute{ctx, cfg}
}

func New(cfg aws.Config) *Compute {
	ctx := context.TODO()
	return &Compute{ctx, cfg}
}
