# yoo_home_server
yoo_home的基于beego框架的服务器API后台
配置说明:
--------------------------------
数据库 --- MySql
使用之前请配置配置文件 --- config/app.confi  
<pre><code>
appname = yoo_home
httpport = 9999  
runmode = dev   
maxmemory = 1<<22    
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
---------------------------
