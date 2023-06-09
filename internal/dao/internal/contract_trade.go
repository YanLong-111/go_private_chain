package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractTradeDao is the data access object for table contract_trade.
type ContractTradeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ContractTradeColumns // columns contains all the column names of Table for convenient usage.
}

// ContractTradeColumns defines and stores column names for table contract_trade.
type ContractTradeColumns struct {
	Id              string // 主键，自增长ID
	TransactionHash string // 交易hash
	ContractAddress string // 合约地址
	UserAddress     string // 用户地址
	TokenId         string // token id
	TokenUri        string // token Uri
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// contractTradeColumns holds the columns for table contract_trade.
var contractTradeColumns = ContractTradeColumns{
	Id:              "id",
	TransactionHash: "transaction_hash",
	ContractAddress: "contract_address",
	UserAddress:     "user_address",
	TokenId:         "token_id",
	TokenUri:        "token_uri",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewContractTradeDao creates and returns a new DAO object for table data access.
func NewContractTradeDao() *ContractTradeDao {
	return &ContractTradeDao{
		group:   "default",
		table:   "contract_trade",
		columns: contractTradeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContractTradeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContractTradeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContractTradeDao) Columns() ContractTradeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContractTradeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContractTradeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContractTradeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
