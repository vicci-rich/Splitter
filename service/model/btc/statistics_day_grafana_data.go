package btc

type StatisticsDayGrafanaData struct {
	ID                           int64 `xorm:"id bigint autoincr pk"`
	Timestamp                    int64 `xorm:"timestamp int notnull unique"`
	ActiveAddressBirthTimeLevel1 int64 `xorm:"active_address_birth_time_level1 bigint notnull default '0'"`
	ActiveAddressBirthTimeLevel2 int64 `xorm:"active_address_birth_time_level2 bigint notnull default '0'"`
	ActiveAddressBirthTimeLevel3 int64 `xorm:"active_address_birth_time_level3 bigint notnull default '0'"`
	ActiveAddressBirthTimeLevel4 int64 `xorm:"active_address_birth_time_level4 bigint notnull default '0'"`
	ActiveAddressBirthTimeLevel5 int64 `xorm:"active_address_birth_time_level5 bigint notnull default '0'"`
	ActiveAddressBirthTimeLevel6 int64 `xorm:"active_address_birth_time_level6 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel1    int64 `xorm:"active_user_birth_time_level1 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel2    int64 `xorm:"active_user_birth_time_level2 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel3    int64 `xorm:"active_user_birth_time_level3 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel4    int64 `xorm:"active_user_birth_time_level4 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel5    int64 `xorm:"active_user_birth_time_level5 bigint notnull default '0'"`
	ActiveUserBirthTimeLevel6    int64 `xorm:"active_user_birth_time_level6 bigint notnull default '0'"`
	ActiveDayCountLevel1         int64 `xorm:"active_day_count_level1 bigint notnull default '0'"`
	ActiveDayCountLevel2         int64 `xorm:"active_day_count_level2 bigint notnull default '0'"`
	ActiveDayCountLevel3         int64 `xorm:"active_day_count_level3 bigint notnull default '0'"`
	ActiveDayCountLevel4         int64 `xorm:"active_day_count_level4 bigint notnull default '0'"`
	ActiveDayCountLevel5         int64 `xorm:"active_day_count_level5 bigint notnull default '0'"`
	ActiveDayCountLevel6         int64 `xorm:"active_day_count_level6 bigint notnull default '0'"`
	ActiveDayCountLevel7         int64 `xorm:"active_day_count_level7 bigint notnull default '0'"`
	ActiveDayCountLevel8         int64 `xorm:"active_day_count_level8 bigint notnull default '0'"`
	ActiveDayCountLevel9         int64 `xorm:"active_day_count_level9 bigint notnull default '0'"`
	ActiveDayCountLevel10        int64 `xorm:"active_day_count_level10 bigint notnull default '0'"`
	ActiveDayCountLevel11        int64 `xorm:"active_day_count_level11 bigint notnull default '0'"`
	ActiveDayCountLevel12        int64 `xorm:"active_day_count_level12 bigint notnull default '0'"`
	ActiveDayCountLevel13        int64 `xorm:"active_day_count_level13 bigint notnull default '0'"`
	ActiveDayCountLevel14        int64 `xorm:"active_day_count_level14 bigint notnull default '0'"`
	ActiveDayCountLevel15        int64 `xorm:"active_day_count_level15 bigint notnull default '0'"`
	AddressTxCountLevel1         int64 `xorm:"address_tx_count_level1 bigint notnull default '0'"`
	AddressTxCountLevel2         int64 `xorm:"address_tx_count_level2 bigint notnull default '0'"`
	AddressTxCountLevel3         int64 `xorm:"address_tx_count_level3 bigint notnull default '0'"`
	UserAddressCountLevel1       int64 `xorm:"user_address_count_level1 bigint notnull default '0'"`
	UserAddressCountLevel2       int64 `xorm:"user_address_count_level2 bigint notnull default '0'"`
	UserAddressCountLevel3       int64 `xorm:"user_address_count_level3 bigint notnull default '0'"`
	UserAddressCountLevel4       int64 `xorm:"user_address_count_level4 bigint notnull default '0'"`
	UTXOLevel1                   int64 `xorm:"utxo_level1 bigint notnull default '0'"`
	UTXOLevel2                   int64 `xorm:"utxo_level2 bigint notnull default '0'"`
	UTXOLevel3                   int64 `xorm:"utxo_level3 bigint notnull default '0'"`
	UTXOLevel4                   int64 `xorm:"utxo_level4 bigint notnull default '0'"`
}

func (t StatisticsDayGrafanaData) TableName() string {
	return tableName("statistics_day_grafana_data")
}
