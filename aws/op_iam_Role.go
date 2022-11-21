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
  ExampleInlinePoliciesV1List = []InlinePolicy{
    {Key: "s3-inline-policy", Value: ExampleS3Policy},
  }
  ExampleInlinePoliciesV2List = []InlinePolicy{
    {Key: "dynamodb-inline-policy", Value: ExampleDynamoDbPolicy},
    {Key: "ses-inline-policy", Value: ExampleSesPolicy},
  }
  ExampleInlinePoliciesV3List = []InlinePolicy{
    {Key: "ses-inline-policy", Value: ExampleSesPolicy},
  }

  ExampleInlinePoliciesList = ExampleInlinePoliciesV1List
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

// Attach and Removes the specified managed policy from the specified role. A role can also
// have inline policies embedded with it. To delete an inline policy, use
// DeleteRolePolicy. For information about policies, see Managed policies and
// inline policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (cmp Compute) AttachInlinePolicies() {
  service := iam.NewFromConfig(cmp.Config)
  getRole, err := service.GetRole(context.Background(), &iam.GetRoleInput{
    RoleName: aws.String(ExampleRoleName),
  })
  if err != nil {
    internal.CheckError(fmt.Sprintf("Role (%s) not found.", ExampleRoleName), err)
  }

  if getRole != nil {
    // fmt.Println(rolePoliciesList)
    currentPolicies := map[string]struct{}{}
    // delete policy that no longer exist
    for _, policy := range ExampleInlinePoliciesList {
      currentPolicies[policy.Key] = struct{}{}
    }

    rolePoliciesList, err := service.ListRolePolicies(context.Background(), &iam.ListRolePoliciesInput{
      RoleName: aws.String(ExampleRoleName),
    })
    internal.CheckError(fmt.Sprintf("Role (%s) inline policy not attached.", ExampleRoleName), err)

    toDeletePolicies := map[string]struct{}{}
    for _, policy := range rolePoliciesList.PolicyNames {
      if _, ok := currentPolicies[policy]; !ok {
        internal.OutputColorizedMessage("yellow", fmt.Sprintf("Inline policy (%s) should not be attached.", policy))
        toDeletePolicies[policy] = struct{}{}
      }
    }

    for policy := range toDeletePolicies {
      _, err := service.DeleteRolePolicy(context.TODO(), &iam.DeleteRolePolicyInput{
        PolicyName: aws.String(policy),
        RoleName: aws.String(ExampleRoleName),
      })
      internal.CheckError(fmt.Sprintf("Inline policy (%s) not detached.", policy), err)
      internal.OutputColorizedMessage("green", fmt.Sprintf("Inline policy detached (%s).", policy))
    }

    // add|update
    for _, policy := range ExampleInlinePoliciesList {
      _, err = service.PutRolePolicy(context.Background(), &iam.PutRolePolicyInput{
        RoleName:       aws.String(ExampleRoleName),
        PolicyName:     aws.String(policy.Key),
        PolicyDocument: aws.String(policy.Value),
      })
      internal.CheckError(fmt.Sprintf("Role (%s) inline policy not attached.", ExampleRoleName), err)
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
