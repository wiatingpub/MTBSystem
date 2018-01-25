#### 采用go-micro开发的电影票在线购票系统

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
- 2、数据库设计(ing)
- 3、服务端开发(ing)
- 4、前端开发(ing)
- 5、联调
- 6、优化

### 如何启动程序：
- 1、 ./ctrl.sh build
- 2、 ./ctrl.sh run
- 3、 ./ctrl.sh init
- 4、 ./ctrl.sh start
- 5、 ./ctrl.sh sql
- 6、 ./ctrl.sh login
- 7、 cd /data/deploy/mtbsystem/
- 8、 bash ./build_local.sh all

