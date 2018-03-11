#!/bin/bash
### chmod 添加supervisor mysqld权限
### conf 复制配置文件以及配置数据库表结构信息

Database="mtbsystem"
echo $1
if [ $1 == "chmod" ]; then
    # 给supervisor.sock mysqld.sock 添加权限
    sudo touch /var/run/supervisor.sock
    sudo touch /var/run/mysqld/mysqld.sock
    sudo chmod 777 /var/run/supervisor.sock
    sudo chmod 777 /var/run/mysqld/mysqld.sock
    sudo service supervisor restart
fi
if [ $1 == "conf" ]; then
    # 将supervisor的conf文件复制到指定目录并且重启supervisorctl
    cp /data/deploy/mtbsystem/dockerbase/supervisor/*conf /etc/supervisor/conf.d/
    # 将conf的文件复制到指定目录
    cp /data/deploy/mtbsystem/dockerbase/conf/redis.conf /etc/redis/6379.conf
    mkdir /etc/consul/
    cp /data/deploy/mtbsystem/dockerbase/conf/consul.json /etc/consul/consul.json

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
fi