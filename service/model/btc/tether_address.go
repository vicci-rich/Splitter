package btc

type TetherAddress struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	Address           string `xorm:"address varchar(255) notnull unique"`
	BirthTimestamp    int64  `xorm:"birth_timestamp int notnull default '0' index"`
	LatestTxTimestamp int64  `xorm:"latest_tx_timestamp int notnull default '0' index"`
	Value             int64  `xorm:"value bigint notnull default '0' index"`
}

func (t TetherAddress) TableName() string {
	return tableName("tether_address")
}
