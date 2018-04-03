# yoo_home_server
yoo_home的基于beego框架的服务器API后台
配置说明:
--------------------------------
数据库 --- MySql
使用之前请配置配置文件 --- config/app.confi  
<pre><code>
appname = yoo_home
httpport = 7000
runmode = dev
maxmemory = 1<<22

host_DB = "127.0.0.1"
port_DB = "3306"
charset = "utf8"
name_DB = "yoo_home"
username_DB = "root"
password_DB = 78901214
  
</code></pre>
安装使用:
---------------------------

服务端运行程序在yoo_home_run文件夹下。直接拷贝yoo_home_run到本地即可运行     

linux下执行nohup命令   
<pre><code>
nohup ./yoo_home_server &  
</code></pre>

windows下运行yoo_home_server.exe文件

接口说明:
--------------------------------                                                                                                                    
接口具体使用方法请用浏览器打开接口测试页面，查看接口测试页面源码                       
例:127.0.0.1:7000/register                                   
 
注册:ip:port/register                      
                              
登录:ip:port/login                                   

相册:ip:port/albums                           

相片:ip:port/photos                                                            

活动:ip:port/activities                                    

活动项:ip:port/activityitems                          

日程安排:ip:port/arrangement                     

附件:ip:port/files                        

验证码:ip:port/captcha                             

用户信息:ip:port/user                                

家庭信息:ip:port/family                               

联系我们:
---------------------------
QQ:78901214  
QQ群:434752626
