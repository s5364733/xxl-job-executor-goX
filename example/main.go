package main

import (
	"github.com/go-basic/ipv4"
	"github.com/s5364733/xxl-job-executor-goX"
	"github.com/s5364733/xxl-job-executor-goX/example/task"
	"log"
)

func main() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://fzxxljob-dev-cstest.fzzqxf.com/xxl-job-admin/"),
		xxl.AccessToken("8OiqaEWzy4PyhQ3V"), //请求令牌(默认为空)
		xxl.ExecutorIp(ipv4.LocalIP()),      //可自动获取
		xxl.ExecutorPort("9999"),            //默认9999（非必填）
		xxl.RegistryKey("golang-jobs"),      //执行器名称
		//xxl.SetLogger(&xxl.Xlogger{}),       //自定义日志
	)
	exec.Init()
	////设置日志查看handler
	//exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
	//	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
	//		FromLineNum: req.FromLineNum,
	//		ToLineNum:   2,
	//		LogContent: "这个是自定义日志handler\n" +
	//			"这个是自定义日志handler\n" +
	//			"这个是自定义日志handler\n" +
	//			"这个是自定义日志handler\n" +
	//			"这个是自定义日志handler\n" +
	//			"这个是自定义日志handler\n",
	//		IsEnd: true,
	//	}}
	//})
	//注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)
	log.Fatal(exec.Run())
}

//// xxl.Logger接口实现
//type logger struct{}
//
//func (l *logger) Info(format string, a ...interface{}) {
//	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
//
//func (l *logger) Error(format string, a ...interface{}) {
//	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
