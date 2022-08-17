terraform {
  backend "s3" {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key    = "log/db_tranaza_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}