#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u docker -pdocker trout_analyzer < "/docker-entrypoint-initdb.d/00_create_tables.sql"