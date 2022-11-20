package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	_ "github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws"

	"awshowto/internal"
)

const (
	ExampleRoleName   = "ik-golang-example-role"
	ExamplePolicyARN  = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
	ExampleSLRService = "elasticbeanstalk.amazonaws.com"
)

var (
	ExampleInlinePoliciesV1 = map[string]string{
		"s3-inline-policy": ExampleS3Policy,
	}
)

func (cmp Compute) CreateRole() {
	service := iam.NewFromConfig(cmp.Config)
	getRole, err := service.GetRole(context.Background(), &iam.GetRoleInput{
		RoleName: aws.String(ExampleRoleName),
	})
	if getRole != nil {
		internal.ExitErrorf(fmt.Sprintf("Role found!!!. role arn: %s.", *getRole.Role.Arn))
	}
	if err != nil {
		// snippet-start:[iam.go-v2.CreateRole]
		// CreateRole
		role, err := service.CreateRole(context.Background(), &iam.CreateRoleInput{
			RoleName:    aws.String(ExampleRoleName),
			Description: aws.String("My super awesome example role"),
			AssumeRolePolicyDocument: aws.String(`{
        "Version": "2012-10-17",
        "Statement": [
        {
          "Sid": "EC2AssumeRole",
          "Effect": "Allow",
          "Principal": {
            "Service": "ec2.amazonaws.com"
          },
            "Action": "sts:AssumeRole"
          }
        ]
      }`),
		})
		if err != nil {
			internal.CheckError("Couldn't create role.", err)
		}
		internal.OutputColorizedMessage("green", "✅ Role Created")
		internal.OutputColorizedMessage("green", fmt.Sprintf("Account:: %s.", *role.Role.Arn))
		// snippet-end:[iam.go-v2.CreateRole]
	}
}

func (cmp Compute) AttachInlinePolicies() {
	service := iam.NewFromConfig(cmp.Config)
	getRole, err := service.GetRole(context.Background(), &iam.GetRoleInput{
		RoleName: aws.String(ExampleRoleName),
	})
	if err != nil {
		internal.CheckError(fmt.Sprintf("Role (%s) not found.", ExampleRoleName), err)
	}
	if getRole != nil {
    for name, doc := range ExampleInlinePoliciesV1 {
        _, err = service.PutRolePolicy(context.Background(), &iam.PutRolePolicyInput{
          RoleName: aws.String(ExampleRoleName),
          PolicyName: aws.String(name),
          PolicyDocument: aws.String(doc),
        })
        if err != nil {
          internal.CheckError(fmt.Sprintf("Role (%s) inline policy not attached.", ExampleRoleName), err)
        }
    }
		internal.OutputColorizedMessage("green", fmt.Sprintf("Inline policies attached to role (%s).", ExampleRoleName))
	}
}

func (cmp Compute) DeleteRole() {
	service := iam.NewFromConfig(cmp.Config)
	getRole, err := service.GetRole(context.Background(), &iam.GetRoleInput{
		RoleName: aws.String(ExampleRoleName),
	})
	if err != nil {
		internal.CheckError(fmt.Sprintf("Role (%s) not found.", ExampleRoleName), err)
	}
	if getRole != nil {
		_, err = service.DeleteRole(context.Background(), &iam.DeleteRoleInput{
			RoleName: aws.String(ExampleRoleName),
		})
		internal.CheckError(fmt.Sprintf("FAILED > Role (%s) deletion.", ExampleRoleName), err)
		internal.OutputColorizedMessage("green", fmt.Sprintf("✅ SUCCESS > Role (%s) deleted.", ExampleRoleName))
	}
}

func (cmp Compute) ListRoles() {
	service := iam.NewFromConfig(cmp.Config)

	roles, err := service.ListRoles(context.Background(), &iam.ListRolesInput{})

	if err != nil {
		internal.CheckError("Could not list roles.", err)
	}
	fmt.Println("🔰 List Roles")
	for _, idxRole := range roles.Roles {
		fmt.Printf("ID: %s. NAME: %s. ARN: %s.",
			*idxRole.RoleId,
			*idxRole.RoleName,
			*idxRole.Arn)
		if idxRole.Description != nil {
			fmt.Printf("DESC: %s", *idxRole.Description)
		}
		fmt.Print("\n")
	}
}
