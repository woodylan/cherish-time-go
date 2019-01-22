# 惜时光后端程序（golang）

## 简介

这是**惜时光**小程序的后端接口，这个项目使用Golang开发，采用了beego+gorm框架。



## 小程序前端项目

惜时光：[https://github.com/woodylan/cherish-time ](https://github.com/woodylan/cherish-time )



## 使用的开源包

- beego：[https://github.com/astaxie/beego](https://github.com/astaxie/beego)
- gorm：[https://github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)
- 微信小程序 SDK：[https://github.com/medivhzhan/weapp](https://github.com/medivhzhan/weapp)




## 安装步骤

1. 安装beego

   ```shell
   go get -u github.com/astaxie/beego
   go get -u github.com/beego/bee
   ```

2. 安装orm

   ```shell
   go get -u github.com/jinzhu/gorm
   ```

3. 安装小程序SDK
 ```shell
   go get -u github.com/medivhzhan/weapp
 ```
4. 导入数据库

   ~~~mysql
   //连接数据库
   mysql -uroot -p

   //新建数据库
   create database cherish_time;

   //选择数据库
   use cherish_time;

   //设置数据库编码
   set names utf8mb4_unicode_ci;

   //导入sql文件
   source /wwwroot/publish/cherish-time.sql;
   ~~~

5. 配置项目

   编辑 `conf/app.conf`

   ```shell
   appname = cherish-time-go
   httpport = 8080 //端口
   runmode = dev
   autorender = false
   copyrequestbody = true
   EnableDocs = true

   # 数据库配置
   db.host =
   db.user =
   db.password =
   db.port = 3306
   db.name = cherish_time
   db.prefix = tb_
   db.timezone = Asia/Shanghai

   # Redis配置
   redis_key = cherishTime
   redis_host = 127.0.0.1
   redis_port = 6379
   redis_password =

   # 微信小程序配置
   weChat_mini_program_appId =
   weChat_mini_program_secret =
   ```

6. 打包发布

   发布到linux

   ```shell
   bee pack -be GOOS=linux
   ```

   会得到一个可执行压缩包，解压到服务器运行，打开8080端口就可以运行了

   ```shell
   tar -zxvf cherish-time-go.tar.gz
   nohup ./cherish-time-go &
   ```

7. 配置nginx

   如果你不想在域名输入端口号，有两个方法，一是用80端口运行，二是做反向代理，这里以nginx做反向代理为例。

   ```nginx
   server {
       listen  80;
       server_name www.example.com;
       root /wwwroot/publish/cherish-time-php/public;
   
       access_log /wwwroot/publish/cherish-time-go/logs/access.log;
       error_log /wwwroot/publish/cherish-time-go/logs/error.log;
       index index.php index.html index.htm;
   
       # ssl ------
       listen       443 ssl;
       ssl_certificate      /etc/letsencrypt/live/example.com/fullchain.pem;
       ssl_certificate_key  /etc/letsencrypt/live/example.com/privkey.pem;
   
       location /api/ {
           try_files /_not_exists_ @backend;
       }
   
       location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           proxy_set_header Host            $http_host;
   
           proxy_pass http://127.0.0.1:8080;
       }
   }
   ```
