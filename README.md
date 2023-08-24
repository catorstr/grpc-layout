# 基于gRPC-Gateway搭建的一个基础项目

> 当前环境依赖:go=1.19.4,make=4.2.1,protoc=3.19.4;依赖环境需自行安装

> 注:接口未增加任何认证

### 命令

```bash
	make init #下载Grpc相关插件
#借用于kratos这个工具的一些便利来加快开发
# 安装kratos 
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
#kratos这个框架本身已经很强大，但是不太喜欢它函数的调用方式(个人习惯，勿喷)，但是kratos这个工具可以合理利用。
#利用kratos初始化项目时，指定自己喜欢的模板：
 	kratos new helloworld -r https://gitee.com/linghuchu-101/grpc-layout.git
#添加proto文件
	kratos proto add api/helloworld/helloworld.proto #生成的proto文件go包指定的位置是'helloworld/proto/helloworld'由根目录开始,建议指定未'./proto/helloworld'不然生成的包在奇奇怪怪的位置。
#编译proto文件:kratos proto client proto/helloworld/helloworld.proto,由于该模板和官方的差别有点大，使用字符命令的时候回报两个警告，看着不太舒服,所以自定义一下。使用如下编译proto文件：(和官方的是有区别的，但不管怎么样，反正都是自动生成的)
    make proto
# 生成service代码 kratos proto server [proto源文件路劲] -t [生成的go文件目标路径]
kratos proto server api/helloworld/helloworld.proto -t handler/service
#好了，还是做不到十全十美，这种方法需要自己在handler/app.go文件中完成服务的注册。emm 官方的也是类似，但是官方用了wire作为自动化注入。

```
