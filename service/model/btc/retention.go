package btc

type Retention struct {
	ID                 int64 `xorm:"id bigint autoincr pk"`
	Timestamp          int64 `xorm:"timestamp bigint notnull index"`
	NewAddressCount    int64 `xorm:"new_address_count bigint notnull"`
	AddressRetention1  int64 `xorm:"address_retention_1_month bigint notnull"`
	AddressRetention2  int64 `xorm:"address_retention_2_month bigint notnull"`
	AddressRetention3  int64 `xorm:"address_retention_3_month bigint notnull"`
	AddressRetention4  int64 `xorm:"address_retention_4_month bigint notnull"`
	AddressRetention5  int64 `xorm:"address_retention_5_month bigint notnull"`
	AddressRetention6  int64 `xorm:"address_retention_6_month bigint notnull"`
	AddressRetention7  int64 `xorm:"address_retention_7_month bigint notnull"`
	AddressRetention8  int64 `xorm:"address_retention_8_month bigint notnull"`
	AddressRetention9  int64 `xorm:"address_retention_9_month bigint notnull"`
	AddressRetention10 int64 `xorm:"address_retention_10_month bigint notnull"`
	AddressRetention11 int64 `xorm:"address_retention_11_month bigint notnull"`
	NewUserCount       int64 `xorm:"new_user_count bigint notnull"`
	UserRetention1     int64 `xorm:"user_retention_1_month bigint notnull"`
	UserRetention2     int64 `xorm:"user_retention_2_month bigint notnull"`
	UserRetention3     int64 `xorm:"user_retention_3_month bigint notnull"`
	UserRetention4     int64 `xorm:"user_retention_4_month bigint notnull"`
	UserRetention5     int64 `xorm:"user_retention_5_month bigint notnull"`
	UserRetention6     int64 `xorm:"user_retention_6_month bigint notnull"`
	UserRetention7     int64 `xorm:"user_retention_7_month bigint notnull"`
	UserRetention8     int64 `xorm:"user_retention_8_month bigint notnull"`
	UserRetention9     int64 `xorm:"user_retention_9_month bigint notnull"`
	UserRetention10    int64 `xorm:"user_retention_10_month bigint notnull"`
	UserRetention11    int64 `xorm:"user_retention_11_month bigint notnull"`
}

func (t Retention) TableName() string {
	return tableName("retention")
}
