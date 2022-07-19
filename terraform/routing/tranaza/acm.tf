resource "aws_acm_certificate" "root" {
  domain_name = data.aws_route53_zone.this.name

  validation_method = "DNS"
  tags = {
    Name = "tranaza-link"
  }

  lifecycle {
    create_before_destroy = true
  }

}

resource "aws_acm_certificate_validation" "root" { // 検証完了でapply
  certificate_arn = aws_acm_certificate.root.arn
}