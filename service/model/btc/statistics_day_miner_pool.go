package btc

type StatisticsDayMinerPool struct {
	ID              int64   `xorm:"id bigint autoincr pk"`
	Name            string  `xorm:"name varchar(100) notnull unique(name_time)"`
	Timestamp       int64   `xorm:"timestamp bigint notnull unique(name_time)"`
	CalcForce       float64 `xorm:"calc_force double notnull default '0'"`
	BlockCount      int64   `xorm:"block_count bigint notnull default '0'"`
	EmptyBlockCount int64   `xorm:"empty_block_count bigint notnull default '0'"`
	SizeAvg         float64 `xorm:"size_avg double notnull default '0'"`
	FeeAvg          float64 `xorm:"fee_avg double notnull default '0'"`
	VOutValueAvg    float64 `xorm:"vout_value_avg double notnull default '0'"`
}

func (t StatisticsDayMinerPool) TableName() string {
	return tableName("statistics_day_miner_pool")
}
