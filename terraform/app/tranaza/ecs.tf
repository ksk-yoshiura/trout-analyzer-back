resource "aws_ecs_cluster" "this" {
  name = "${local.name_prefix}-${local.service_name}"
  
  capacity_providers = [
    "FARGATE", "FARGATE_SPOT"
  ]

  tags = {
    name = "${local.name_prefix}-${local.service_name}"
  }
}

resource "aws_ecs_task_definition" "this" {
  family = "${local.name_prefix}-${local.service_name}"

  task_role_arn = aws_iam_role.ecs_task.arn

  network_mode = "awsvpc"

  requires_compatibilities = [
    "FARGATE",
  ]

  execution_role_arn = aws_iam_role.ecs_task_execution.arn

  memory = "512"
  cpu = "256"

  container_definitions = jsonencode(
    [
      {
        name = "nginx"
        image = "${module.nginx.ecr_repository_this_repository_url}:latest"

        portMappings = [
          {
            containerPort = 80
            protocol = "tcp"
          }
        ]

        dependsOn = [
          {
            containerName = "golang"
            condition = "START"
          }
        ]

        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-group = "/ecs/${local.name_prefix}-${local.service_name}/nginx"
            awslogs-region = data.aws_region.current.id
            awslogs-stream-prefix = "ecs"
          }
        }
      }, 
      {
        name = "golang"
        image = "${module.golang.ecr_repository_this_repository_url}:latest"

        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-group = "/ecs/${local.name_prefix}-${local.service_name}/golang"
            awslogs-region = data.aws_region.current.id
            awslogs-stream-prefix = "ecs"
          }
        }
      }
    ]
  )

  # volume {
  #   name = ""
  # }

  tags = {
    Name = "${local.name_prefix}-${local.service_name}"
  }
}