module "application" {
    source = "./modules/application"

    application_name = "messager-web-app"
    application_port = 8080

    application_environment_variables = [
        {
            name  = "AWS_SQS_QUEUE_URL"
            value = "https://sqs.eu-west-1.amazonaws.com/010526271920/messager-web-app-queue"
            description = "Queue to store the messages"
        },
        {
            name  = "APP_PORT"
            value = "8080"
            description = "Application Port"
        },
        {
            name = "NEW_RELIC_LICENSE_KEY"
            value = "NEW_RELIC_LICENSE_KEY"
            description = "New Relic License Key"
        }
    ]

    allowed_client_ip_list = [
        "0.0.0.0/0"
    ]
    
    allowed_internet_ip_list = [
        "0.0.0.0/0"
    ]

    application_repository = "messager-web-app"
    application_image_tag = "latest"
}