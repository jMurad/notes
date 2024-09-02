#!/bin/bash

CONTAINER_NAME="postgres"
DB_USER="postgres"
DB_PASSWORD="123456"

docker exec -it $CONTAINER_NAME psql -U $DB_USER -c "CREATE DATABASE notes_test;"
