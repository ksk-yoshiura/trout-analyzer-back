resource "aws_eip" "nat_gateway" {
  for_each = var.enable_nat_gateway ? local.nat_gateway_azs : {}
  vpc      = true
  tags = {
    Name = "${aws_vpc.this.tags.Name}-nat-gateway-${each.key}"
  }
}

locals { // コスト戦略
  nat_gateway_azs = var.single_nat_gateway ? { keys(var.azs)[0] = values(var.azs)[0] } : var.azs
}