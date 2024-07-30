# go-net
常用网络相关的组件实现
在工作中常用的通信协议在这里实现一遍，方便以后使用

初步规划包括： http、grpc、jsonrpc
http：常用的框架是gin, 目前整理了一些 demo

grpc：grpc 是目前 golang 中最常用的 rpc 框架，我使用最多的web 框架是 kratos，同时支持 grpc 和 http，通过 proto 来定义接口， 使用了 grpc-gateway 来实现 grpc 的 http 转发，会给上一些 demo。

jsonrpc：jsonrpc 是一个比较简单的 rpc 框架，使用了 jsonrpc 来实现 rpc 通信，在 ethereum 中使用较多。
