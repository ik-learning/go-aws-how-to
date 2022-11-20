package aws

// Contains information about an IAM policy, including the policy document. This
// data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
type InlinePolicy struct {
  // The name of the policy.
	Key string
  // The policy document.
	Value string
}

const (
	ExampleS3Policy = `{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
                "s3:*"
            ],
            "Effect": "Allow",
            "Resource": ["*"],
            "Sid": "S3Allow"
          },
          {
            "Action": [
                "s3:ListBucket",
                "s3:ListBucketVersions",
                "s3:GetBucketLocation",
                "s3:ListBucketMultiPartUploads"
            ],
            "Effect": "Allow",
            "Resource": "*",
            "Sid": "S3AllowList"
          }
        ]
      }`
  ExampleDynamoDbPolicy = `{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
                "dynamodb:*"
            ],
            "Effect": "Allow",
            "Resource": ["*"],
            "Sid": "DynamoDBAllow"
          }
        ]
      }`
  ExampleSesPolicy = `{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
              "ses:SendEmail",
              "ses:SendRawEmail"
            ],
            "Effect": "Allow",
            "Resource": ["*"],
            "Sid": "SESAllow"
          }
        ]
      }`
)
