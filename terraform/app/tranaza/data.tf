data "aws_caller_identity" "self" {}

data "aws_region" "current" {}

data "terraform_remote_state" "network_main" {
  backend = "s3"

  config = {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key    = "network/main_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}

data "terraform_remote_state" "routing_tranaza_link" {
  backend = "s3"

  config = {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key    = "routing/tranaza_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }

}