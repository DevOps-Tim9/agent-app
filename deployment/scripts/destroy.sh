#!/bin/sh

cd terraform || exit
DATABASE_URL=$(heroku config:get DATABASE_URL --app "$TERRAFORM_BACKEND") && export DATABASE_URL
terraform init -backend-config="conn_str=$DATABASE_URL"
terraform destroy -auto-approve -var servers_app_name="$APP_NAME" -var gateway_app_name="${GATEWAY_NAME}"