terraform {
  backend "s3" {
    bucket = "tfstate-s3-bucket-for-tranaza"
    key    = "cicd/tranaza_v1.2.4.tfstate"
    region = "ap-northeast-1"
  }
}