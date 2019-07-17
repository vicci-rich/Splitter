package btc

type StatisticsMonth struct {
	ID                 int64 `xorm:"id bigint autoincr pk"`
	Timestamp          int64 `xorm:"timestamp int notnull unique"`
	ActiveAddressCount int64 `xorm:"active_address_count int notnull default '0'"`
	ActiveUserCount    int64 `xorm:"active_user_count int notnull default '0'"`
}

func (t StatisticsMonth) TableName() string {
	return tableName("statistics_month")
}
