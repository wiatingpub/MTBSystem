#### 采用go-micro开发的电影票在线购票系统

-------------------

系列博客：
- https://mp.weixin.qq.com/s/5bn5ZkAJYR0IEaa5H0bsFg
- https://mp.weixin.qq.com/s/SQ9HS4wKSz8HtNXHOA5oeg
- https://mp.weixin.qq.com/s/Y55hfVF4a8A6XOI5OHHlgw
- https://mp.weixin.qq.com/s/Yo2f-XtbbxI6jrYDTtKxKA

-------------------
#### 模块划分：
![模块划分.png](http://upload-images.jianshu.io/upload_images/3365849-dfaec3d3a064fd8a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 服务划分：
![服务划分.png](http://upload-images.jianshu.io/upload_images/3365849-005e52ef50e643ae.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 数据库ER图
![数据库ER关系图.png](http://upload-images.jianshu.io/upload_images/3365849-9c1abcd5fedd1043.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 技术方案：
- 服务端：go-micro
- 数据库：mysql
- 缓    存：redis
- 前   端：el & vue
- 服务器： centos 7 & nginx
- 本地环境：go1.9
- 容器：Docker
- 进程管理：supervisor
- 数据库备份：冷备份(rsync+mysqldump)

#### 开发进程：
- 1、搭建好开发框架 (get)
- 2、数据库设计(get)
- 3、服务端开发(get)
- 4、前端开发(get)
- 5、联调(get)
- 6、优化(get)

### 如何启动程序：
- 1、 ./ctrl.sh build #构建docker环境，构建完成后可以省略该步骤
- 2、 ./ctrl.sh run #启动docker容器环境
- 3、 ./ctrl.sh init conf #环境配置，包括数据库 
- 4、 ./ctrl.sh init chmod  #权限设定
- 5、 ./ctrl.sh start  #启动容器
- 6、 ./ctrl.sh login #登录容器
- 7、 cd /data/deploy/mtbsystem/
- 8、 bash ./build_local.sh api-srv #启动api服务
- 9、 bash ./build_local.sh all #启动所有服务

### 如何添加服务
- 1、 在proto下添加文件，如cms.ext.proto
- 2、 在src下添加cms-srv
- 3、 在dockerbase/supervisor下添加cms-srv-conf
- 4、 ./ctrl.sh init conf
- 5、 ./ctrl.sh login
- 6、 cd /data/deploy/mtbsystem/
- 7、 bash ./build_local.sh cms-rv

### mysql冷备份
- 1、 启动： bash mysql_backup.sh
- 2、 数据恢复：gzip -d mtbsystem-xxxx.sql.gz
- 3、 数据回复：mysql -u username -p database < 文件名 

### 效果演示
- 1、前台访问(手机网站)：http://front.lixifan.cn/
- 2、后台访问:http://admin.lixifan.cn/#/login admin 123456 / 新光影城 xgyc 

-------

**Java源码分析、go语言应用、微服务，更多干货欢迎关注公众号：**

![公众号.jpg](https://user-gold-cdn.xitu.io/2019/5/29/16aff62fa6acf090?w=258&h=258&f=jpeg&s=16132)


