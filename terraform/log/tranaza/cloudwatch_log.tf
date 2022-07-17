resource "aws_cloudwatch_log_group" "nginx" {
  name = "/ecs/${local.service_name}-${local.env_name}-nginx"

  retention_in_days = 90
}

resource "aws_cloudwatch_log_group" "golang" {
  name = "/ecs/${local.service_name}-${local.env_name}-golang"

  retention_in_days = 90
}