# Get AWS Regions
data "aws_region" "current" {}

# Get Account ID
data "aws_caller_identity" "current" {}

# Get Default VPC
data "aws_vpc" "default" {}

# Get Default Subnets
data "aws_subnets" "private" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}
