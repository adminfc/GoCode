package main
/* 
Go modules是golang 1.11新特性，用来管理萨夫尼克的包的依赖关系。

使用方法(前两个最常用)：

初始化模块
go mod init <项目模块名称>
依赖关系处理，根据go.mod文件
go mod tidy
将依赖包复制到项目下的vendor目录
go mod vendor   //如果包被屏蔽，可以使用这个命令，随后使用go build -mod=vendor编译
显示依赖关系
go list -m all
显示 详细依赖关系
go list -m -json all
下载依赖
go mod download [path@version] //中括号内容非必需

远程包：
在go官网找到包，然后里面有命令go get -u github地址,然后再import package

示例：
1.先在项目文件夹下面go mod init 文件夹名
2. cd到包文件夹下面，用go build创建缓存？
3. import package
4. go mod tidy
*/
func main() {

}
