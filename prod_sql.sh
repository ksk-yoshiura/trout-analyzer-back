#!/bin/bash -e
HOGE=hoge
echo $HOGE
mysql -h tranaza-prod-tranaza.csbzfseof6ef.ap-northeast-1.rds.amazonaws.com --port 3306 -u tranaza -p $DB_PASSWORD  -e < "./db/mysql_init/00_create_tables.sql"