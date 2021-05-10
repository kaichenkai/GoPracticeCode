package main

//Go语言的标准库覆盖网络、系统、加密、编码、图形等各个方面，可以直接使用标准库的 http 包进行 HTTP 协议的收发处理；
//网络库基于高性能的操作系统通信模型（Linux 的 epoll、Windows 的 IOCP）；所有的加密、编码都内建支持，不需要再从第三方开发者处获取

//Go 语言的编译器也是标准库的一部分，通过词法器扫描源码，使用语法树获得源码逻辑分支等。Go 语言的周边工具也是建立在这些标准库上。在标准库上可以完成几乎大部分的需求

//Go 语言的标准库以包的方式提供支持，下表是 Go 语言标准库中常见的包及其功能

/*
Go语言标准库包名		功能
bufio				带缓冲的 I/O 操作
bytes				实现字节操作
container			封装堆、列表和环形列表等容器
crypto				加密算法
database			数据库驱动和接口
debug				各种调试文件格式访问及调试功能
encoding			常见算法如 JSON、XML、Base64 等
flag				命令行解析
fmt					格式化操作
go					Go 语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改
html				HTML 转义及模板系统
image				常见图形格式的访问及生成
io					实现 I/O 原始访问接口及访问封装
math				数学库
net					网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等
os					操作系统平台不依赖平台操作封装
path				兼容各操作系统的路径操作实用函数
plugin				Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载
reflect				语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值
regexp				正则表达式封装
runtime				运行时接口
sort				排序接口
strings				字符串转换、解析及实用函数
time				时间接口
text				文本模板及 Token 词法器
*/