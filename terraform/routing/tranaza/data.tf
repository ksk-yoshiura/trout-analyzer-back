data "terraform_remote_state" "network_main" {
  backend = "s3"

  config = {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key = "network/main_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}

data "terraform_remote_state" "log_alb" {
  backend = "s3"

  config = {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key = "log/alb_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}