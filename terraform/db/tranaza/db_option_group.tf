resource "aws_db_option_group" "this" {
  name = "${local.service_name}-${local.env_name}-${local.service_name}"

  engine_name          = "mysql"
  major_engine_version = "8.0"

  tags = {
    Name = "${local.service_name}-${local.env_name}-${local.service_name}"
  }
}