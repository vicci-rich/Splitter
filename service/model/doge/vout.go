package doge

type VOut struct {
	ID                 int64  `xorm:"id bigint autoincr pk"`
	TxID               string `xorm:"tx_id char(64) notnull unique(tx_id_number)"`
	BlockHeight        int64  `xorm:"block_height int notnull index(block_height)"`
	Value              uint64 `xorm:"value bigint notnull"`
	Address            string `xorm:"address varchar(256) notnull index(address)"`
	Timestamp          int64  `xorm:"timestamp int notnull index(timestamp)"`
	ScriptPublicKey    string `xorm:"script_pubkey text notnull"`
	Type               string `xorm:"type varchar(256) notnull"`
	RequiredSignatures int64  `xorm:"required_signatures smallint notnull"`
	Number             int64  `xorm:"number smallint notnull unique(tx_id_number)"`
	IsUsed             int64  `xorm:"is_used tinyint notnull"`
	IsCoinbase         int64  `xorm:"is_coinbase tinyint notnull"`
}

func (t VOut) TableName() string {
	return tableName("vout")
}
