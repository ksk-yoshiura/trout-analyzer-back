#!/bin/bash
DB_PASSWORD=(`aws ssm get-parameters --names "/tranaza/DB_PASSWORD" --query "Parameters[*].{Value: Value}" --with-decryption --output text`)

mysql -h tranaza-prod-tranaza.csbzfseof6ef.ap-northeast-1.rds.amazonaws.com --port 3306 -u tranaza -p${DB_PASSWORD}
# aws ecs execute-command --cluster tranaza-prod-tranaza --task f8263624a62d49dbabd4d6a97744fbb0 --container golang --interactive --command 