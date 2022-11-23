resource "aws_s3_bucket" "env_file" { // envファイル用
  bucket = "tfstate-s3-bucket-for-tranaza-${local.name_prefix}-${local.service_name}-env-file"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  tags = {
    Name = "tfstate-s3-bucket-for-tranaza-${local.name_prefix}-${local.service_name}-env-file"
  }
}

resource "aws_s3_bucket" "public_image" { // 画像ファイル
  bucket = "${local.name_prefix}-${local.service_name}-s3-image-upload"
  acl    = "public-read"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  cors_rule {
    allowed_origins = ["https://tranaza.com"] 
    allowed_methods = ["GET"]
    allowed_headers = ["*"]
    max_age_seconds = 3000
  }

  tags = {
    Name = "${local.name_prefix}-${local.service_name}-s3-image-uploade"
  }

}