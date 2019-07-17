package doge

type Address struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	Address           string `xorm:"address varchar(255) notnull unique(address)"`
	BirthTimestamp    int64  `xorm:"birth_timestamp int notnull default '0' index(birth_time)"`
	LatestTxTimestamp int64  `xorm:"latest_tx_timestamp int notnull default '0'"`
	Value             int64  `xorm:"value bigint notnull default '0' index(value)"`
}

func (t Address) TableName() string {
	return tableName("address")
}
