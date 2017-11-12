#!/bin/bash

# 创建数据库和表
mysql -u root -e "CREATE DATABASE mtbsystem"
R=/data/deploy/mtbsystem/sql/
cd $R
for var in `ls`
do
    bash $var
done

# 更新服务
cd ..
bash build_local.sh all

