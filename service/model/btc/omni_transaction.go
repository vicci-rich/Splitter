package btc

type OmniTansaction struct {
	ID               int64  `xorm:"id bigint autoincr pk"`
	TxID             string `xorm:"tx_id char(64) notnull unique(tx_id)"`
	SendingAddress   string `xorm:"sending_address varchar(64) notnull index(sending_address)"`
	ReferenceAddress string `xorm:"reference_address varchar(64) notnull index(reference_address)"`
	BlockHeight      int64  `xorm:"block_height bigint notnull index(block_height)"`
	Timestamp        int64  `xorm:"timestamp bigint notnull index(timestamp)"`
	Version          int64  `xorm:"version bigint notnull"`
	TypeInt          int64  `xorm:"type_int  bigint notnull"`
	Type             string `xorm:"type varchar(64) notnull"`
	PropertyID       int64  `xorm:"propertyid bigint notnull index(propertyid)"`
	Amount           uint64 `xorm:"amount bigint notnull"`
	Number           int64  `xorm:"number bigint notnull"`
	Fee              uint64 `xorm:"fee bigint notnull"`
	Valid            int64  `xorm:"valid bigint notnull"`
}

func (t OmniTansaction) TableName() string {
	return tableName("omni_transaction")
}
