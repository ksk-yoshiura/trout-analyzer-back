resource "aws_ecs_cluster" "this" {
  name = "${local.name_prefix}-${local.service_name}"
  
  capacity_providers = [
    "FARGATE", "FARGATE_SPOT"
  ]

  tags = {
    name = "${local.name_prefix}-${local.service_name}"
  }
}
