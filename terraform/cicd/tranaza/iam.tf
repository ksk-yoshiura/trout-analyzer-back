resource "aws_iam_user" "github" {
  name = "tranaza-prod-github"

  tags = {
    Name = "tranaza-prod-github"
  }
}

resource "aws_iam_role" "deployer" {
  name = "tranaza-prod-deployer"

  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "sts:AssumeRole",
            "sts:TagSession"
          ],
          "Principal": {
            "AWS": aws_iam_user.github.arn
          }
        }
      ]
    }
  )

  tags = {
    Name = "tranaza-prod-deployer"
  }
}

data "aws_iam_policy" "ecr_power_user" {
  arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
}

resource "aws_iam_role_policy_attachment" "role_deployer_policy_ecr_power_user" {
  role = aws_iam_role.deployer.name
  policy_arn = data.aws_iam_policy.ecr_power_user.arn
}
