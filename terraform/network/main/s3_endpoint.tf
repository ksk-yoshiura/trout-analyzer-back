// 画像アップロード用S3
resource "aws_vpc_endpoint" "s3_endpoint" {
  vpc_id            = aws_vpc.this.id
  service_name      = "com.amazonaws.${data.aws_region.current.name}.s3"
  vpc_endpoint_type = "Gateway"

  tags = {
    Name = "${aws_vpc.this.tags.Name}"
  }
}
