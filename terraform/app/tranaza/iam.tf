resource "aws_iam_role" "ecs_task_execution" {
  name = "${local.service_name}-${local.env_name}-ecs-task-execution"

  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "ecs-tasks.amazonaws.com"
          },
          "Action" : "sts:AssumeRole"
        }
      ]
    }
  )

  tags = {
    Name = "${local.service_name}-${local.env_name}-ecs-task-execution"
  }
}

data "aws_iam_policy" "ecs_task_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = data.aws_iam_policy.ecs_task_execution.arn
}


resource "aws_iam_role" "ecs_task" {
  name = "${local.name_prefix}-${local.service_name}-ecs-task"

  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "ecs-tasks.amazonaws.com"
          },
          "Action" : "sts:AssumeRole"
        }
      ]
    }
  )
  tags = {
    Name = "${local.name_prefix}-${local.service_name}-ecs-task"
  }
}

resource "aws_iam_policy" "ssm" {
  name = "${local.name_prefix}-${local.service_name}-ssm"
  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "ssm:GetParameters",
            "ssm:GetParameter"
          ],
          "Resource" : "arn:aws:ssm:${data.aws_region.current.id}:${data.aws_caller_identity.self.account_id}:parameter/${local.service_name}/*"
        }
      ]
    }
  )
  tags = {
    Name = "${local.name_prefix}-${local.service_name}-ssm"
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_ssm" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = aws_iam_policy.ssm.arn
}

// AWS SDK GOでGetParameterするため
resource "aws_iam_role_policy_attachment" "ecs_task_ssm" {
  role       = aws_iam_role.ecs_task.name
  policy_arn = aws_iam_policy.ssm.arn
}


resource "aws_iam_role_policy" "ecs_task_ssm" { // ECS EXECのため
  name = "ssm"
  role = aws_iam_role.ecs_task.id

  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "ssmmessages:CreateControlChannel",
            "ssmmessages:CreateDataChannel",
            "ssmmessages:OpenControlChannel",
            "ssmmessages:OpenDataChannel",
            "s3:*", // なぜかこれでアップロードできるようになった。なんで？？
            "iam:GetGroup",
            "ec2:*",
          ],
          "Resource" : "*"
        }
      ]
    }
  )
}

resource "aws_iam_policy" "s3_env_file" { // envファイルアクセスのため
  name = "${local.name_prefix}-${local.service_name}-s3-env-file"
  policy = jsonencode(
    {
      "Version" : "2012-10-17"
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : "s3:GetObject"
          "Resource" : "${aws_s3_bucket.env_file.arn}/*"
        },
        {
          "Effect" : "Allow",
          "Action" : "s3:GetBucketLocation",
          "Resource" : aws_s3_bucket.env_file.arn
        }
      ]
    }
  )
  tags = {
    Name = "${local.name_prefix}-${local.service_name}-s3-env-file"
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_s3_env_file" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = aws_iam_policy.s3_env_file.arn
}

resource "aws_iam_policy" "s3_image_upload" { // 画像ファイルアップロード
  name = "${local.name_prefix}-${local.service_name}-s3-image-upload"
  policy = jsonencode(
    {
      "Version" : "2012-10-17"
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "s3:GetObject",
            "s3:ListBucket",
            "s3:PutObject",
            "s3:PutObjectAcl"
          ],
          "Resource" : [
            aws_s3_bucket.public_image.arn,
            "${aws_s3_bucket.public_image.arn}/*"
          ]
        }
      ]
    }
  )
  tags = {
    Name = "${local.name_prefix}-${local.service_name}-s3-image-upload"
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_s3_image_upload" {
  role       = aws_iam_role.ecs_task.name
  policy_arn = aws_iam_policy.s3_image_upload.arn
}
