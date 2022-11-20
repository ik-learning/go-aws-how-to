package aws

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
)
