package bch

type StatisticsDayOHLC struct {
	ID        int64   `xorm:"id bigint autoincr pk"`
	Timestamp int64   `xorm:"timestamp int notnull unique"`
	Open      float64 `xorm:"open decimal(38,4) notnull"`
	High      float64 `xorm:"high decimal(38,4) notnull"`
	Low       float64 `xorm:"low decimal(38,4) notnull"`
	Close     float64 `xorm:"close decimal(38,4) notnull"`
}

func (t StatisticsDayOHLC) TableName() string {
	return tableName("statistics_day_ohlc")
}
