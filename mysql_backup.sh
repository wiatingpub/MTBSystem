#!/bin/bash
# 以下配置信息请自己修改
mysql_user="root" #MySQL备份用户
mysql_password="" #MySQL备份用户的密码
mysql_host="localhost"
mysql_port="3306"
mysql_charset="utf8" #MySQL编码
backup_db_arr=("mtbsystem") #要备份的数据库名称，多个用空格分开隔开 如("db1""db2" "db3")
backup_location=backup/mysql  #备份数据存放位置，末尾请不要带"/",此项可以保持默认，程序会自动创建文件夹
expire_backup_delete="OFF" #是否开启过期备份删除 ON为开启 OFF为关闭
expire_days=3 #过期时间天数 默认为三天，此项只有在expire_backup_delete开启时有效
# 以下不需要修改
backup_time=`date +%Y%m%d%H%M`  #定义备份详细时间
backup_Ymd=`date +%Y-%m-%d` #定义备份目录中的年月日时间
backup_3ago=`date -d '3 days ago' +%Y-%m-%d` #3天之前的日期
backup_dir=$backup_location/$backup_Ymd  #备份文件夹全路径
welcome_msg="Welcome to use MySQL backup tools!" #欢迎语
# 判断MYSQL是否启动,mysql没有启动则备份退出
mysql_ps=`ps -ef | grep mysql | wc -l`
mysql_listen=`netstat -an |grep LISTEN |grep $mysql_port|wc -l`
if [ [ $mysql_ps == 0 ] -o [$mysql_listen == 0 ] ]; then
	echo "ERROR:MySQL is not running! backup stop!"
	exit
else
	echo $welcome_msg
fi
# 连接到mysql数据库，无法连接则备份退出
mysql -h$mysql_host -P$mysql_port -u$mysql_user -p$mysql_password <<end
	use mysql;
	select host,user from user where user='root' and host='localhost';
	exit
end
flag=`echo $?`
if [ $flag != "0" ]; then
	echo "ERROR:Can't connect mysql server! backup stop!"
	exit
else
	echo "MySQL connect ok! Please wait......"
	# 判断有没有定义备份的数据库，如果定义则开始备份，否则退出备份
	if [ "$backup_db_arr" != "" ];then
	    #dbnames=$(cut -d ',' -f1-5 $backup_database)
		#echo "arr is (${backup_db_arr[@]})"
		for dbname in ${backup_db_arr[@]}
		do
			echo "database $dbname backup start..."
			`mkdir -p $backup_dir`
			`mysqldump  -h$mysql_host  -P$mysql_port  -u$mysql_user -p$mysql_password $dbname --default-character-set=$mysql_charset | gzip> $backup_dir/$dbname-$backup_time.sql.gz` flag=`echo $?`
			if [ $flag == "0" ];then
				echo "database $dbname success backup to $bac
				kup_dir/$dbname-$backup_time.sql.gz"
			else
				echo "database $dbname backup fail!"
			fi
		done
	else
		echo "ERROR:No database to backup! backup stop"
		exit
	fi
	# 如果开启了删除过期备份，则进行删除操作
	if [ "$expire_backup_delete" == "ON" -a  "$backup_location" != "" ];then
                #`find $backup_location/ -type d -o -type f -ctime +$expire_days -exec rm -rf {} \;`
	`find $backup_location/ -type d -mtime +$expire_days | xargs
    rm -rf`
            echo "Expired backup data delete complete!"
            fi
            echo "All database backup success! Thank you!"
            exit
    fi