output "subnets" {
    description = "Subnets for the application"
    value = data.aws_subnets.private.ids
}

output "vpc_id" {
    description = "VPC ID for the application"
    value = data.aws_vpc.default.id
}

output "service_account_name" {
    description = "Service Account Name for the application"
    value = module.application_service_account.iam_role_name
}

output "service_account_arn" {
    description = "Service Account ARN for the application"
    value = module.application_service_account.iam_role_arn
}

output "apprunner_service_id" {
    description = "App Runner Service ID for the application"
    value = module.app_runner.service_id
}