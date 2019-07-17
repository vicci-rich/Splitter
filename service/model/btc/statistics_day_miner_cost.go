package btc

type StatisticsDayMinerCost struct {
	ID            int64   `xorm:"id bigint autoincr pk"`
	Timestamp     int64   `xorm:"timestamp int notnull unique"`
	Cost          float64 `xorm:"cost double notnull default '0'"`
	BTCOutput     float64 `xorm:"btc_output double notnull default '0'"`
	EnergyCPerDay float64 `xorm:"energy_c_per_day double notnull default '0'"`
	EnergyCPerBtc float64 `xorm:"energy_c_per_btc double notnull default '0'"`
	Income        float64 `xorm:"income double notnull default '0'"`
	Profit        float64 `xorm:"profit double notnull default '0'"`
}

func (t StatisticsDayMinerCost) TableName() string {
	return tableName("statistics_day_miner_cost")
}
