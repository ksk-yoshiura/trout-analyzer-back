#!/bin/bash
set -eu

$DB_PASSWORD = =(`aws ssm get-parameters --names /tranaza/DB_PASSWORD --query "Parameters[*].{Value: Value}"  --output text`)

mysql -h tranaza-prod-tranaza.csbzfseof6ef.ap-northeast-1.rds.amazonaws.com --port 3306 -u tranaza -p $DB_PASSWORD  -eu < "./db/mysql_init/00_create_tables.sql"