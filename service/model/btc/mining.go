package btc

type Mining struct {
	ID            int64  `xorm:"id bigint autoincr pk"`
	Address       string `xorm:"address varchar(255) notnull unique"`
	CoinbaseTimes int64  `xorm:"coinbase_times int notnull default '0'"`
	PoolName      string `xorm:"pool_name varchar(255)"`
}

func (t Mining) TableName() string {
	return tableName("mining")
}
