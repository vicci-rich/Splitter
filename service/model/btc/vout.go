package btc

type VOut struct {
	ID                 int64  `xorm:"id bigint autoincr pk"`
	TxID               string `xorm:"tx_id char(64) notnull"`
	BlockHeight        int64  `xorm:"block_height int notnull"`
	Value              uint64 `xorm:"value bigint notnull"`
	Address            string `xorm:"address varchar(256) notnull"`
	Timestamp          int64  `xorm:"timestamp int notnull"`
	ScriptPublicKey    string `xorm:"script_pubkey text notnull"`
	Type               string `xorm:"type varchar(256) notnull"`
	RequiredSignatures int64  `xorm:"required_signatures smallint notnull"`
	Number             int64  `xorm:"number smallint notnull"`
	IsUsed             int64  `xorm:"is_used tinyint notnull"`
	IsCoinbase         int64  `xorm:"is_coinbase tinyint notnull"`
}

func (t VOut) TableName() string {
	return tableName("vout")
}
