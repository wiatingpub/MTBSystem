#!/bin/sh
Database="mtbsystem"

# 创建数据库和表
mysql -u root -e "CREATE DATABASE $Database"
R=/data/deploy/mtbsystem/sql/
cd $R
for var in `ls`
do
    if [[ "${var}" != "data" ]]; then
           mysql -u root -D "$Database" -e "source $var"
    fi
done

cd data
for var in `ls`
do
    mysql -u root -D "$Database" -e "source $var"
done
