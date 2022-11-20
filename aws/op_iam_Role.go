package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	_ "github.com/aws/aws-sdk-go-v2/service/sts"

	"awshowto/internal"
)

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
