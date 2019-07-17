package eth

type MinerPoolAddress struct {
	ID      int64  `xorm:"id bigint autoincr pk"`
	Name    string `xorm:"name varchar(150) notnull"`
	Address string `xorm:"address char(40) notnull"`
}

func (t MinerPoolAddress) TableName() string {
	return tableName("miner_pool_address")
}
