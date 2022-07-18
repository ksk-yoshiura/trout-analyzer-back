module "nginx" {
  source = "../../modules/ecr"

  name = "${local.service_name}-${local.env_name}-nginx"
}

module "golang" {
  source = "../../modules/ecr"

  name = "${local.service_name}-${local.env_name}-golang"
}