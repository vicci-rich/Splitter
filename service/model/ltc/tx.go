package ltc

type Transaction struct {
	ID          int64  `xorm:"id bigint autoincr pk"`
	TxID        string `xorm:"tx_id char(64) notnull unique(tx_id_block_height)"`
	BlockHeight int64  `xorm:"block_height int notnull unique(tx_id_block_height)"`
	Timestamp   int64  `xorm:"timestamp int notnull"`
	Version     int64  `xorm:"version int notnull"`
	Size        int64  `xorm:"size int notnull"`
	VSize       int64  `xorm:"vsize int notnull"`
	LockTime    int64  `xorm:"lock_time bigint notnull"`
	Hash        string `xorm:"hash char(64) notnull"`
	Weight      int64  `xorm:"weight int notnull"`
	Number      int64  `xorm:"number int notnull"`
	VInCount    int    `xorm:"vin_count smallint notnull"`
	VInValue    int64  `xorm:"vin_value bigint notnull"`
	VOutCount   int    `xorm:"vout_count smallint notnull"`
	VOutValue   uint64 `xorm:"vout_value bigint notnull"`
	Fee         int64  `xorm:"fee bigint notnull"`
}

func (t Transaction) TableName() string {
	return tableName("transaction")
}
