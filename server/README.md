### Chat-test Server
* * * * 
#### 介绍(Introduction)
这是Chat-test 的Server
##### 功能(Function)

* 监听tcp port，等待client连接。**future:连接管理，同一个用户顶号，重连等问题**
* client 加入一个全局聊天室。 保存client名字和登录时间。**future:聊天室管理，根据聊天室id,token验证加入**
* client加入聊天室后，获取最近的历史记录（通过循环队列保存历史记录）
* 动态保存最近时间内最流行的词汇 
* 使用ac自动机过滤敏感词汇，同时有处理噪音词

##### 目录(catalog)
chat:聊天室
examples:简单用例
network:server网络
service:处理client请求

#### 使用(Usages)
* * * *
exmaples 有简单的使用用例  
启动成功显示：  
time="2020-09-18T01:36:24+08:00" level=info msg="注册消息: MsgID=ReqLoginE"  
time="2020-09-18T01:36:24+08:00" level=info msg="注册消息: MsgID=ReqChatE"  
time="2020-09-18T01:36:24+08:00" level=info msg="注册消息: MsgID=ReqHeartbeatE"  
time="2020-09-18T01:36:24+08:00" level=info msg="加载网络Profanity word, https://raw.githubusercontent.com/RobertJGabriel/Google-profanity-words/master/list.txt"   component=filter  
time="2020-09-18T01:36:25+08:00" level=info msg="Profanity word：4r5e" component=filter  
time="2020-09-18T01:36:25+08:00" level=info msg="Profanity word：5h1t" component=filter  
time="2020-09-18T01:36:25+08:00" level=info msg="Profanity word：5hit" component=filter  
time="2020-09-18T01:36:25+08:00" level=info msg="Profanity word：a55" component=filter  
time="2020-09-18T01:36:25+08:00" level=info msg=BuildFailurePointer.... component=filter  
"......"    
time="2020-09-18T01:36:25+08:00" level=info msg="profanity word 加载完毕" component=room  
time="2020-09-18T01:36:25+08:00" level=info msg="聊天房间, Init成功" component=room  
time="2020-09-18T01:36:25+08:00" level=warning msg="Chat服务器启动成功, TCP监听地址=:660, 心跳间隔=5秒 ,1600364185542802900"  

Notes: 有时候会获取网络上 profanity words 失败  
time="2020-09-18T02:10:38+08:00" level=info msg="加载网络Profanity word, https://raw.githubusercontent.com/RobertJGabriel/Google-profanity-words/master/list.txt" component=filter  
panic: Get "https://raw.githubusercontent.com/RobertJGabriel/Google-profanity-words/master/list.txt": wsarecv: An existing connection was forcibly closed by the remote host.

为了方便测试  
so: 添加获取本地 profanity words 文件的接口 