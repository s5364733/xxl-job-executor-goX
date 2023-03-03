package xxl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
用来日志查询，显示到xxl-job-admin后台
*/

type LogHandler func(req *LogReq) *LogRes

// 默认返回
func defaultLogHandler(req *LogReq) *LogRes {
	filename := fmt.Sprintf("log_%d", req.LogID)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Errorf("err %v", err)
	}

	fd, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("read to fd fail", err)
	}

	return &LogRes{Code: SuccessCode, Msg: "", Content: LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  string(fd),
		IsEnd:       true,
	}}
}

// 请求错误
func reqErrLogHandler(w http.ResponseWriter, req *LogReq, err error) {
	res := &LogRes{Code: FailureCode, Msg: err.Error(), Content: LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   0,
		LogContent:  err.Error(),
		IsEnd:       true,
	}}
	str, _ := json.Marshal(res)
	_, _ = w.Write(str)
}
