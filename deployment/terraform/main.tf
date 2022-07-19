terraform {
  backend "pg" {
  }
}

provider "heroku" {
}

variable "servers_app_name" {
  description = "Unique name of the Servers app"
}

variable "gateway_app_name" {
  description = "Unique name of the Gateway app"
}

resource "heroku_config" "servers" {
  vars = {
    STAGE = "PROD"
    CORS_ORIGIN = "https://${heroku_app.gateway.name}.herokuapp.com"
  }
}

resource "heroku_app" "servers" {
  name   = var.servers_app_name
  region = "eu"
  stack  = "container"
}

resource "heroku_build" "servers" {
  app_id = heroku_app.servers.id

  source {
    path = "servers"
  }
}

resource "heroku_app_config_association" "servers" {
  app_id = heroku_app.servers.id

  vars = heroku_config.servers.vars
}

resource "heroku_addon" "database" {
  app_id  = heroku_app.servers.id
  plan = "heroku-postgresql:hobby-dev"
}

resource "heroku_config" "gateway" {
  vars = {
    API_URL = "${heroku_app.servers.name}.herokuapp.com"
  }
}

resource "heroku_app" "gateway" {
  name   = var.gateway_app_name
  region = "eu"
  stack  = "container"
}

resource "heroku_app_config_association" "gateway" {
  app_id = heroku_app.gateway.id

  vars = heroku_config.gateway.vars
}

resource "heroku_build" "gateway" {
  app_id = heroku_app.gateway.id

  source {
    path = "gateway"
  }
  depends_on = [
    null_resource.gateway_build
  ]
}

data "template_file" "gateway_build" {
  template = file("${path.module}/gateway/heroku.tpl")
  vars = {
    api_url = "\\\"  apiUrl:  'https://${heroku_app.servers.name}.herokuapp.com/api/server'\\\""
  }
}

resource "null_resource" "gateway_build" {
  triggers = {
    template = data.template_file.gateway_build.rendered
  }

  provisioner "local-exec" {
    command = "echo \"${data.template_file.gateway_build.rendered}\" > ${path.module}/gateway/heroku.yml"
  }
}

output "servers_app_url" {
  value = "https://${heroku_app.servers.name}.herokuapp.com"
}
output "gateway_app_url" {
  value = "https://${heroku_app.gateway.name}.herokuapp.com"
}