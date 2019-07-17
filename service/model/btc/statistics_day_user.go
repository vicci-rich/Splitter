package btc

type StatisticsDayUser struct {
	ID                   int64 `xorm:"id bigint autoincr pk"`
	Timestamp            int64 `xorm:"timestamp int notnull unique"`
	UserCount            int64 `xorm:"user_count bigint notnull default '0'"`
	UserNew              int64 `xorm:"user_new bigint notnull default '0'"`
	UserActive           int64 `xorm:"user_active bigint notnull default '0'"`
	ExchangeVInValueSum  int64 `xorm:"exchange_vin_value_sum bigint notnull default '0'"`
	ExchangeVOutValueSum int64 `xorm:"exchange_vout_value_sum bigint notnull default '0'"`
}

func (t StatisticsDayUser) TableName() string {
	return tableName("statistics_day_user")
}
