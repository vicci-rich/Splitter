package eth

type StatisticsWeek struct {
	ID                      int64 `xorm:"id bigint autoincr pk"`
	Timestamp               int64 `xorm:"timestamp int notnull unique index"`
	ActiveAddressCount      int64 `xorm:"active_address_count decimal(38,0) notnull default '0'"`
	ActiveTokenAddressCount int64 `xorm:"active_token_address_count decimal(38,0) notnull default '0'"`
}

func (t StatisticsWeek) TableName() string {
	return tableName("statistics_week")
}
