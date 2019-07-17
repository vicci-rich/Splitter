package eth

type StatisticsMonthTokens struct {
	ID                        int64  `xorm:"id bigint autoincr pk"`
	Timestamp                 int64  `xorm:"timestamp int notnull default '0' index"`
	TokenAddress              string `xorm:"token_address char(40) notnull default ''"`
	TokenTxActiveAddressCount int    `xorm:"token_tx_active_address_count int notnull default '0'"`
}

func (t StatisticsMonthTokens) TableName() string {
	return tableName("statistics_month_tokens")
}
