#!/bin/bash -e
mysql -h tranaza-prod-tranaza.csbzfseof6ef.ap-northeast-1.rds.amazonaws.com --port 3306 -u tranaza -p $DB_PASSWORD trout_analyzer < "./db/mysql_init/00_create_tables.sql"