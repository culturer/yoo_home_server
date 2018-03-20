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
=============================   
注册:
测试页面 --- ip:port/register 例 : 127.0.0.1:7000/register
接口 --- 请求类型:POST
         请求参数: "password" --- 密码                
                   "tel" --- 电话号码                  
          返回参数:
        <pre><code>   
            {
                 "msg": " 注册失败,该手机号已被注册 ",
                 "status": 400,
                 "time": "2018-03-20 11:13:54"
            }

            {
                 "msg": "register success ",
                 "status": 200,
                 "time": "2018-03-20 11:14:48"
            }
        </code></pre>
        
联系我们:
---------------------------
QQ:78901214  
QQ群:434752626
