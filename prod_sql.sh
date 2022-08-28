#!/bin/bash
echo $DB_PASSWORD
#mysql -h tranaza-prod-tranaza.csbzfseof6ef.ap-northeast-1.rds.amazonaws.com --port 3306 -u tranaza -p $DB_PASSWORD -e trout_analyzer < "./db/00_create_tables.sql"