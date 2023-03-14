package v1

import "github.com/gogf/gf/v2/frame/g"

type NewJobTaskReq struct {
	g.Meta     `path:"/user/new-job-task" method:"post" tags:"JobService" summary:"上传一个新的工作任务"`
	Ciphertext string `v:"required"`
}

type NewJobTaskRes struct{}
