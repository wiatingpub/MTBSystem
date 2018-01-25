#!/bin/bash

# 将supervisor的conf文件复制到指定目录并且重启supervisorctl
cp /data/deploy/mtbsystem/dockerbase/supervisor/*conf /etc/supervisor/conf.d/
# 将conf的文件复制到指定目录
cp /data/deploy/mtbsystem/dockerbase/conf/redis.conf /etc/redis/6379.conf
mkdir /etc/consul/
cp /data/deploy/mtbsystem/dockerbase/conf/consul.json /etc/consul/consul.json

# 给supervisor.sock mysqld.sock 添加权限
sudo touch /var/run/supervisor.sock
sudo touch /var/run/mysqld/mysqld.sock
sudo chmod 777 /var/run/supervisor.sock
sudo chmod 777 /var/run/mysqld/mysqld.sock
sudo service supervisor restart
