package doge

type Mining struct {
	ID            int64  `xorm:"id bigint autoincr pk"`
	Address       string `xorm:"address varchar(255) notnull unique(address)"`
	CoinbaseTimes int64  `xorm:"coinbase_times int notnull default '0'"`
	PoolName      string `xorm:"pool_name varchar(256)"`
}

func (t Mining) TableName() string {
	return tableName("mining")
}
