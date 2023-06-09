package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserData is the golang structure of table user_data for DAO operations like Where/Data.
type UserData struct {
	g.Meta        `orm:"table:user_data, do:true"`
	Id            interface{} //
	UserId        interface{} // 用户id
	Opcode        interface{} // 唯一操作码
	UserNick      interface{} // 用户昵称
	AccountHash   interface{} // 用户创建时hash
	UserAddress   interface{} // 用户地址
	ParentId      interface{} // 父级信息id
	CurrentStatus interface{} // 状态信息 0:任务已加入 1:任务已完成
}
