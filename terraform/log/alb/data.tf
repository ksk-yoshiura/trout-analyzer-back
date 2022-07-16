data "aws_elb_service_account" "current" {}

output "s3_bucket_this_id" {
  value = aws_s3_bucket.this.id
}