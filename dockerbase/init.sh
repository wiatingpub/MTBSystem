#!/bin/bash
Database="mtbsystem"

# 将supervisor的conf文件复制到指定目录并且重启supervisorctl
cp /data/deploy/mtbsystem/dockerbase/supervisor/*conf /etc/supervisor/conf.d/
# 将conf的文件复制到指定目录
cp /data/deploy/mtbsystem/dockerbase/conf/redis.conf /etc/redis/6379.conf
mkdir /etc/consul/
cp /data/deploy/mtbsystem/dockerbase/conf/consul.json /etc/consul/consul.json
# 重新用supervisor加载进程
supervisorctl reload
# 创建数据库和表
mysql -u root -e "CREATE DATABASE $Database"
R=/data/deploy/mtbsystem/sql/
cd $R
for var in `ls`
do
   mysql -u root -D "$Database" -e "source $var"
done

# 更新服务
cd ..
bash build_local.sh all

