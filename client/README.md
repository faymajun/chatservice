### Chat-test Client
* * * * 
#### 介绍(Introduction)
这是Chat-test 的client
##### 功能(Features)
* 连接聊天室服务器，加入聊天室。Notes:需要先启动server。**future:还没有做断线重连等处理。**
* 给聊天室发送聊天消息
* 通过命令 "/popular" 获取最流行词汇
* 通过命令 "/stats bob" 获取bob的登录时间
##### 目录(catalog)
examples:简单用例
network:client网络
service:处理服务器返回消息
* * * * 
#### 使用(Usages)
exmaples 有简单的使用用例
* * * * 
启动成功显示：  
time="2020-09-18T03:43:51+08:00" level=info msg="注册消息: MsgID=ResLoginE"  
time="2020-09-18T03:43:51+08:00" level=info msg="注册消息: MsgID=ResChatE"  
time="2020-09-18T03:43:51+08:00" level=info msg="注册消息: MsgID=ResHistoryChatE"  
time="2020-09-18T03:43:51+08:00" level=info msg="注册消息: MsgID=ResHeartbeatE"  
time="2020-09-18T03:43:52+08:00" level=warning msg="Chat Client启动成功, 服务器地址=:660, 心跳间隔=5秒 ,1600371832113245000"  
time="2020-09-18T03:43:52+08:00" level=info msg="历史聊天消息长度：  39" component=client  
// 此处省略历史内容  
time="2020-09-18T03:43:52+08:00" level=info msg="name:\"bob\" content:\"****boy 222 \" " component=client  
time="2020-09-18T03:43:52+08:00" level=info msg="name:\"bob\" content:\"****o ***** world 333\" " component=client  
time="2020-09-18T03:43:52+08:00" level=info msg="name:\"bob\" content:\"****o ***** world 444\" " component=client  
time="2020-09-18T03:43:57+08:00" level=info msg="content:\"world\" " component=client      // **popular**  
time="2020-09-18T03:43:57+08:00" level=info msg="content:\"00d 00h 00m 08s\" " component=client  //**请求stats**  
