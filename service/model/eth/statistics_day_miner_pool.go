package eth

type StatisticsDayMinerPool struct {
	ID                 int64   `xorm:"id bigint autoincr pk"`
	Timestamp          int64   `xorm:"timestamp int notnull"`
	Name               string  `xorm:"name varchar(100) notnull"`
	HashRate           float64 `xorm:"hash_rate decimal(38,4) notnull default '0'"`
	BlockCount         int64   `xorm:"block_count int notnull default '0'"`
	UncleCount         int64   `xorm:"uncle_count int notnull default '0'"`
	EmptyBlockCount    int64   `xorm:"empty_block_count int notnull default '0'"`
	BlockSizeAvg       float64 `xorm:"block_size_avg decimal(38,4) notnull default '0'"`
	UncleSizeAvg       float64 `xorm:"uncle_size_avg decimal(38,4) notnull default '0'"`
	FeeAvg             float64 `xorm:"fee_avg decimal(38,4) notnull default '0'"`
	BlockRewardAvg     float64 `xorm:"block_reward_avg decimal(38,4) notnull default '0'"`
	ReferenceRewardAvg float64 `xorm:"reference_reward_avg decimal(38,4) notnull default '0'"`
	UncleRewardAvg     float64 `xorm:"uncle_reward_avg decimal(38,4) notnull default '0'"`
}

func (t StatisticsDayMinerPool) TableName() string {
	return tableName("statistics_day_miner_pool")
}
