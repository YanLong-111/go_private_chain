package entity

// UserData is the golang structure for table user_data.
type UserData struct {
	Id            int    `json:"id"             description:""`
	UserId        string `json:"user_id"        description:"用户id"`
	UserAddress   string `json:"user_address"   description:"用户地址"`
	ParentId      int    `json:"parent_id"      description:"父级信息id"`
	CurrentStatus int    `json:"current_status" description:"状态信息 0:任务已加入 1:任务已完成"`
}
