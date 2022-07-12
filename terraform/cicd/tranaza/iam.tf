resource "aws_iam_user" "github" {
  name = "tranaza-prod-github"

  tags = {
    Name = "tranaza-prod-github"
  }
}

# resource "aws_iam_role" "deployer" {
#   name = "tranaza-prod-deployer"

#   assume_role_policy = jsonencode(
#     {
#       "Version" : "2012-10-17",
#       "Statement" : [
#         {
#           "Effect" : "Allow",
#           "Action" : [
#             "sts:AssumeRole",
#             "sts:TagSession"
#           ],
#           "Pricipal": {
#             "AWS": aws_iam_user.github.arn
#           }
#         }
#       ]
#     }
#   )

#   tags = {
#     Name = "${locals.name_prefix}-${locals.service_name}-deployer"
#   }
# }

