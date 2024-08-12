# Luminor test task

## Prerequisites 

- Terraform S3 Bucket created for states

## Getting started

1. Clone the repository

```bash
git clone git@github.com:KostLinux/luminor-platform-engineering-test-task.git
```

2. Configure ECR Repo and SQS Queue via cloudformation

**Note:** Replace 123456789012 with your AWS Account ID. Additionally ParameterValues which you're setting needs to be configured in terraform as well.

```bash
aws cloudformation create-stack --stack-name luminor-test-task-preparation --template-body file://cloudformation/prepare-platform.yml --parameters ParameterKey=Repos
itoryName,ParameterValue=messager-web-app ParameterKey=QueueName,ParameterValue=messager-web-app-queue --capabilities CAPABILITY_NAMED_IAM --region eu-west-1
```

3. Build and push the Docker image

```bash
docker build -t messager-web-app .
docker tag messager-web-app:latest 123456789012.dkr.ecr.eu-west-1.amazonaws.com/messager-web-app:latest
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 123456789012.dkr.ecr.eu-west-1.amazonaws.com
docker push 123456789012.dkr.ecr.eu-west-1.amazonaws.com/messager-web-app:latest
```

4. Deploy the application

```bash
cd terraform
terraform init
terraform apply
```


## Troubleshooting

If you will get:
```bash
╷
│ Error: creating App Runner Service (messager-web-app): operation error AppRunner: CreateService, https response error StatusCode: 400, RequestID: 41b25720-9939-487b-90c1-9dda93340b24, InvalidRequestException: Error in assuming access role arn:aws:iam::010526271920:role/aws-app-runner-role
│ 
│   with module.application.module.app_runner.aws_apprunner_service.this[0],
│   on .terraform/modules/application.app_runner/main.tf line 34, in resource "aws_apprunner_service" "this":
│   34: resource "aws_apprunner_service" "this" {
│ 
╵
```

Rerun terraform apply. There's an issue with IAM Role module maintained by AntonBabenko. It's not a part of the task to fix it.

## Comments from Author

- There might be some connectivity issues with SQS. No time this week to investigate it further.
- The task didn't commit to use secret stores like AWS SSM, Hashicorp Vault, Github Secrets etc, so i didn't use them.
- The task didn't commit to use CI/CD, so i didn't use it.