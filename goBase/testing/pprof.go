package main

//使用方式
//1. 导入
//import (
//    _ "net/http/pprof"
//)
//
//2. 访问 "http://localhost:8080/debug/pprof/"

//3. 性能数据的剖析
//（1）命令行交互方式
//$go tool pprof xxx.test cpu.prof // 剖析通过性能基准测试采集的数据
//$go tool pprof standalone_app cpu.prof // 剖析独立程序输出的性能采集数据
//通过net/http/pprof注册的性能采集数据服务端点获取数据并剖析
//$go tool pprof http://localhost:8080/debug/pprof/profile

//（2）Web图形化方式
