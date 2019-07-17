package btc

type StatisticsDayTransaction struct {
	ID                       int64   `xorm:"id bigint autoincr pk"`
	Timestamp                int64   `xorm:"timestamp bigint notnull unique"`
	ValueDayConsume          float64 `xorm:"value_day_consume double notnull default '0'"`
	TxRate                   float64 `xorm:"tx_rate double notnull default '0'"`
	TxCount                  int64   `xorm:"tx_count bigint notnull default '0'"`
	TxVInAvg                 float64 `xorm:"tx_vin_avg double notnull default '0'"`
	TxVOutAvg                float64 `xorm:"tx_vout_avg double notnull default '0'"`
	TxSizeAvg                float64 `xorm:"tx_size_avg double notnull default '0'"`
	TxValueSum               float64 `xorm:"tx_value_sum double notnull default '0'"`
	TxValueAvg               float64 `xorm:"tx_value_avg double notnull default '0'"`
	TxFeeAvg                 float64 `xorm:"tx_fee_avg double notnull default '0'"`
	TxSizeFeeAvg             float64 `xorm:"tx_size_fee_avg double notnull default '0'"`
	TxValueWithoutHotAddress float64 `xorm:"tx_value_without_hot_address double notnull default '0'"`
	TxValueWithoutLongChain  float64 `xorm:"tx_value_without_long_chain double notnull default '0'"`
	BlockTradeCount          int64   `xorm:"block_trade_count bigint notnull default '0'"`
	BlockTradeSum            int64   `xorm:"block_trade_sum bigint notnull default '0'"`
	AddressCountSum          int64   `xorm:"address_count_sum bigint notnull default '0'"`
	AddressCountNew          int64   `xorm:"address_count_new bigint notnull default '0'"`
	ActiveAddressCount       int64   `xorm:"active_address_count bigint notnull default '0'"`
	AddressVInVOutAvg        float64 `xorm:"address_vin_vout_avg double notnull default '0'"`
	AddressTxValueAvg        float64 `xorm:"address_tx_value_avg double notnull default '0'"`
	ActivePercent            float64 `xorm:"active_percent double notnull default '0'"`
	StoreRate                float64 `xorm:"store_rate double notnull default '0'"`
	RatioOfMarketValue       float64 `xorm:"ratio_of_market_value double notnull default '0'"`
	Freshness                float64 `xorm:"freshness double notnull default '0'"`
	DormantAddress           int64   `xorm:"dormant_address bigint notnull default '0'"`
	DeadAddress              int64   `xorm:"dead_address bigint notnull default '0'"`
	WakeUpAddress            int64   `xorm:"wake_up_address bigint notnull default '0'"`
	RebornAddress            int64   `xorm:"reborn_address bigint notnull default '0'"`
}

func (t StatisticsDayTransaction) TableName() string {
	return tableName("statistics_day_transaction")
}
