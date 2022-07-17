module "docker" {
  source = "../../modules/ecr"

  name = "${local.service_name}-${local.env_name}-docker"
}

module "golang" {
  source = "../../modules/ecr"

  name = "${local.service_name}-${local.env_name}-golang"
}