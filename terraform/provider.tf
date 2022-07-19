provider "aws" {
  region = "ap-northeast-1"

  default_tags {
    tags = {
      Env    = "prod"
      System = "tranaza"

    }
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.42.0"
    }
  }

  required_version = "1.2.4"

}