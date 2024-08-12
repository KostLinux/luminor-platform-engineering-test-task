terraform {
  required_version = ">= 0.12"
  backend "s3" {
    bucket         = "terraform-sandbox-s3"
    key            = "luminor-test-task.tfstate"
    region         = "eu-west-1"
    encrypt        = true
  }
}