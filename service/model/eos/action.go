package eos

type Action struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNum        int64  `xorm:"block_num int index"`
	Account         string `xorm:"account char(12) notnull"`
	Name            string `xorm:"name varchar(20) notnull"`
	Authorization   string `xorm:"authorization varchar(128) notnull"`
	Data            string `xorm:"data text notnull"`
	HexData         string `xorm:"hex_data text"`
}

func (t Action) TableName() string {
	return tableName("action")
}
