module "docker" {
  source = "../../modules/ecr"

  name = "tranaza-prod-docker"
}

module "golang" {
  source = "../../modules/ecr"

  name = "${locals.service_name}-${locals.env_name}-golang"
}