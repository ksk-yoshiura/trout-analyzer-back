#!/bin/bash
set -eu

DB_PASSWORD=(`aws ssm get-parameters --names "/tranaza/DB_PASSWORD" --query "Parameters[*].{Value: Value}" --with-decryption --output text`)

mysql -h "${DB_HOST}"  --port 3306 -u tranaza -p"${DB_PASSWORD}" -eu < "./db/mysql_init/00_create_tables.sql"