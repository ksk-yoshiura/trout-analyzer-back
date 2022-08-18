data "aws_kms_alias" "rds" {
  name = "alias/aws/rds"
}

data "terraform_remote_state" "network_main" {
  backend = "s3"

  config = {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key    = "network/main_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}