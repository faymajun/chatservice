### go_modules
* * *
这是 chat-test 的 go 实现  
Notes:  
如果获取依赖包出错 like this:  
golang代理超时报错"https://proxy.golang.org/github.com/********** timeout  
解决方法只需要换一个国内能访问的代理即可，终端执行以下命令  
go env -w GOPROXY=https://goproxy.cn