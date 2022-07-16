data "aws_route53_zone" "this" {
  name = "trout-analyzer.com" // 取得済み
}

resource "aws_route53_zone" "certificate_validation" {
  for_each = {
    for dvo in aws_acm_certification.root.domain_validation_options : dvo.domain_name => {
      name = dvo.resource_record_name
      type = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  name = each.value.name
  records = [each.value.record]
  ttl = 60
  type = each.value.type
  zone_id = data.aws_route53_zone.this.id
}