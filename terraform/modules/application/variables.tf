variable "allowed_client_ip_list" {
    type    = list(string)
    default = []
}

variable "allowed_internet_ip_list" {
    type    = list(string)
    default = ["0.0.0.0/0"]
}

variable "application_name" {
    description = "Name of the App Runner service"
    type        = string
}

variable "application_port" {
    description = "Port on which the application listens"
    type        = number
}

variable "application_environment_variables" {
    description = "Environment variables for the application"
    type        = list(object({
        name  = string
        value = string
        description = string
    }))
    default     = []
}

variable "application_repository" {
    description = "ECR repository for the application"
    type        = string
}

variable "application_image_tag" {
    description = "Tag for the application image"
    type        = string
}

variable "auto_scaling_configurations" {
    description = "Map of auto scaling configurations"
    type = map(object({
        name            = string
        min_size        = number
        max_size        = number
        max_concurrency = number
    }))
    default = {
        app_runner_autoscaler = {
        name            = "apprunner_autoscaler"
        max_concurrency = 100
        max_size        = 8
        min_size        = 1
        }
    }
}

variable "instance_configuration" {
    description = "Map of instance configuration"
    type = object({
        cpu    = number
        memory = number
    })
    default = {
        cpu    = 1024
        memory = 2048
    }
}