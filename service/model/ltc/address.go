package ltc

type Address struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	Address           string `xorm:"address varchar(255) notnull unique(IDX_ltc_address_address)"`
	BirthTimestamp    int64  `xorm:"birth_timestamp int notnull default '0' index(IDX_ltc_address_birth_time)"`
	LatestTxTimestamp int64  `xorm:"latest_tx_timestamp int notnull default '0'"`
	Value             int64  `xorm:"value bigint notnull default '0' index(IDX_ltc_address_value)"`
}

func (t Address) TableName() string {
	return tableName("address")
}
