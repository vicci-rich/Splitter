package eth

type Token struct {
	ID            int64  `xorm:"id bigint autoincr pk"`
	TokenAddress  string `xorm:"token_address char(40) notnull unique index"`
	DecimalLength int64  `xorm:"decimal_len int notnull"`
	Name          string `xorm:"name varchar(128) notnull"`
	Symbol        string `xorm:"symbol varchar(128) notnull"`
	TotalSupply   string `xorm:"total_supply varchar(128) null"` // follows attribute temerally can't obtain
	Owner         string `xorm:"owner char(40) notnull"`
	Timestamp     int64  `xorm:"timestamp int notnull index"`
}

func (t Token) TableName() string {
	return tableName("token")
}
