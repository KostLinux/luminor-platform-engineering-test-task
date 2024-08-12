resource "aws_security_group" "app_runner" {
    name = "app_runner"
    description = "AppRunner Security Group"
    vpc_id = data.aws_vpc.default.id
    tags = {
        Name = "AppRunner Security Group"
        Type = "Private"
    }

    ingress {
        description = "Allow traffic from the load balancer"
        from_port   = var.application_port
        to_port     = var.application_port
        protocol    = "tcp"
        cidr_blocks = sort(distinct(var.allowed_client_ip_list))
    }

    egress {
        description = "Allow all outbound traffic"
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_blocks = sort(distinct(var.allowed_internet_ip_list))
    }
}