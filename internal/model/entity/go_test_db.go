package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoTestDb is the golang structure for table go_test_db.
type GoTestDb struct {
	Id              int         `json:"id"              description:"自增ID"`
	Opcode          string      `json:"opcode"          description:"opcode"`
	ContractName    string      `json:"contractName"    description:"contract name"`
	ContractAddress string      `json:"contractAddress" description:"contract address"`
	ContractHash    string      `json:"contractHash"    description:"contract hash"`
	GasUsed         int64       `json:"gasUsed"         description:"gas 使用量"`
	GasUsdt         float64     `json:"gasUsdt"         description:"消耗的gas转为usdt"`
	ChainId         int64       `json:"chainId"         description:"链id"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"合约创建时间"`
	CurrentStatus   int         `json:"currentStatus"   description:"合约创建状态 0:任务提交 1:任务进行中 2:任务完成"`
}
