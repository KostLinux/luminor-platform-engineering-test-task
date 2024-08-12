# Create IAM Service Account that can access S3 buckets, connect to database
#tfsec:ignore:aws-iam-no-policy-wildcards
module "iam_policy" {
  source = "terraform-aws-modules/iam/aws//modules/iam-policy"
  version = "~> 5.39.0"

  name        = "aws-app-runner-policy"
  path        = "/"
  description = "Policy for AWS App Runner Access to Amazon services"

  policy = <<-EOF
  {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Action": [
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:DescribeImages",
                "ecr:GetAuthorizationToken",
                "ecr:BatchCheckLayerAvailability"
              ],
              "Resource": [
                  "*"
              ]                
          },
          {
            "Effect": "Allow",
            "Action": [
              "apprunner:*"
            ],
            "Resource": "*"
          },
          {
            "Effect": "Allow",
            "Action": [
              "logs:CreateLogGroup",
              "logs:CreateLogStream",
              "logs:PutLogEvents"
            ],
            "Resource": "*"
          },
          {
            "Effect": "Allow",
            "Action": [
                "sqs:SendMessage",
                "sqs:ReceiveMessage",
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes",
                "sqs:GetQueueUrl"
            ],
            "Resource": "*"
          }
      ]
  }
  EOF

  tags = {
    ManagedBy = "Terraform"
  }

}

module "application_service_account" {
  source        = "terraform-aws-modules/iam/aws//modules/iam-assumable-role"
  version = "~> 5.39.0"

  role_name          = "aws-app-runner-role"
  
  custom_role_policy_arns = [
    module.iam_policy.arn
  ]

  trusted_role_services = [
    "apprunner.amazonaws.com",
    "build.apprunner.amazonaws.com",
  ]

  create_role = true
  
  role_requires_mfa = false

  depends_on = [ module.iam_policy ]

  tags = {
    ManagedBy = "Terraform"
  }

}