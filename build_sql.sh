#!/bin/sh
Database="mtbsystem"

mysql -u root -e "CREATE DATABASE $Database"

R=sql
cd $R
for var in `ls`
do
    if [[ "${var}" != "data" ]]
     then
            mysql -u root -D "$Database" -e "source $var"ba
    fi
done

cd data
for var in `ls`
do
    mysql -u root -D "$Database" -e "source $var"
done

