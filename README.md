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
--------------------------------------------        
测试页面 --- ip:port/register 例 : 127.0.0.1:7000/register                  
接口 --- 请求类型:POST  
        URL: ip:port/register 例 : 127.0.0.1:7000/register                    
        请求参数: "password" --- 密码                 
                   "tel" --- 电话号码                        
        返回参数:                                        
        注册失败                                                            
        <pre><code>   
            {                            
                 "msg": " 注册失败,该手机号已被注册 ",                    
                 "status": 400,                  
                 "time": "2018-03-20 11:13:54"                 
            }                      
        </code></pre>   
        注册成功
        <pre><code>                               
            {                
                 "msg": "register success ",                                
                 "status": 200,                       
                 "time": "2018-03-20 11:14:48"                               
            }                           
         </code></pre>                           
                              
                              
登录:
--------------------------------------------
测试页面 --- ip:port/login 例 : 127.0.0.1:7000/login                  
接口 --- 请求类型:POST  
        URL: ip:port/login 例 : 127.0.0.1:7000/login                    
        请求参数: "password" --- 密码                 
                   "tel" --- 电话号码                        
        返回参数: 
        登录失败                                                            
        <pre><code>   
            {
                "message": " 登陆失败,账号不存在或账号,密码错误。 ",
                "status": 400,
                "time": "2018-03-20 11:23:43"
            }                      
        </code></pre>   
        登录成功
        <pre><code>                               
            {
                "family": {
                    "Id": 2,
                    "FamilyName": "",
                    "FamilyNotifyTitle": " 欢迎来到优家 ",
                    "FamilyNotifyContent": "",
                    "FamilyNotifyTime": "2018-03-18 13:30:18",
                    "CreatedTime": "2018-03-16 17:37:55",
                    "Msg": ""
                },
                "familyUsers": [
                    {
                      "Id": 2,
                      "Username": "",
                      "Password": "123456",
                      "RealName": "",
                      "Sex": false,
                      "Uid": "",
                      "Birth": "",
                      "Tel": "123456",
                      "Email": "",
                      "Icon": "",
                      "NMsg": "",
                      "RelationId": 0,
                      "FamilyId": 2,
                      "FamilyName": "",
                      "FatherId": 0,
                      "MotherId": 0,
                      "MateId": 0,
                      "CreatedTime": "2018-03-16 17:37:55",
                      "LoginTime": "",
                      "Msg": "",
                      "Permission": 0
                    }
                ],
                "message": "login success ",
                "status": 200,
                "time": "2018-03-20 11:21:28",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMyNDQwODgsInVzZXJJZCI6Mn0.SyX3hiU3op7DWEvyyhJvAyG1kBVbOcL14RaGyg5UpNs",
                "user": {
                    "Id": 2,
                    "Username": "",
                    "Password": "123456",
                    "RealName": "",
                    "Sex": false,
                    "Uid": "",
                    "Birth": "",
                    "Tel": "123456",
                    "Email": "",
                    "Icon": "",
                    "NMsg": "",
                    "RelationId": 0,
                    "FamilyId": 2,
                    "FamilyName": "",
                    "FatherId": 0,
                    "MotherId": 0,
                    "MateId": 0,
                    "CreatedTime": "2018-03-16 17:37:55",
                    "LoginTime": "",
                    "Msg": "",
                    "Permission": 0
               }
            }                           
         </code></pre>                 
                                  
                                  
                                    
相册接口:
--------------------------------------------
测试页面 --- ip:port/albums 例 : 127.0.0.1:7000/albums                  
接口 --- 请求类型:POST    
        URL: ip:port/albums 例 : 127.0.0.1:7000/albums                    
        请求参数: 
        options:操作类型                   
                 [options == 0  查询]            
                 [options == 1  增加]             
                 [options == 2  删除]             
                 [options == 3  修改]            
        options == 0            
                albumType:类型             
                [albumType == false  FamilyAlbum]家庭相册                       
                  -需要传入参数userId                  
                [albumType == true  UserAlbum]私人相册                  
                  -需要传入参数familyId               
        返回值:          
        <code><pre>      
        {
              "albums": [
                {
                  "Id": 2,
                  "UserId": 1,
                  "FamilyId": 0,
                  "Name": "我的相册",
                  "Icon": "",
                  "CreateTime": "2018-03-21 14:16:30"
                },
                {
                  "Id": 63,
                  "UserId": 1,
                  "FamilyId": 0,
                  "Name": "测试数据",
                  "Icon": "",
                  "CreateTime": "2018-04-03 10:30:42"
                }
              ],
              "status": 200,
              "time": "2018-04-03 16:44:18"
            }
        </code></pre>                                            
        options == 1 ，新增相册                                
        传入参数:albumType,userId(or familyId),albumItemName                
        返回参数                                  
       <code><pre>                
       {                    
          "albumItemId": 65,
          "status": 200,
          "time": "2018-04-03 16:48:49"
        }                        
       </code></pre>                                       
        options == 2,删除相册                              
        传入参数:albumItemId                 
        返回参数:                
       <code><pre>                
       {
          "message": "delAlbumItem success",
          "status": 200,
          "time": "2018-04-03 16:53:10"
        }                       
       </code></pre>                             
       options == 3,修改相册                 
       传入参数:albumItemId，albumItemName                                  
       返回参数:                  
       <code><pre>                
       {
          "message": "updateAlbumItem success",
          "status": 200,
          "time": "2018-04-03 16:56:08"
        }                       
       </code></pre>                                
       
联系我们:
---------------------------
QQ:78901214  
QQ群:434752626
