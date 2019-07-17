package tron

type Block struct {
	ID               int64  `xorm:"id bigint autoincr pk"`
	BlockNumber      int64  `xorm:"block_number int notnull unique index"`
	BlockHash        string `xorm:"block_hash char(64) notnull index"`
	ParentHash       string `xorm:"previous char(64) notnull"`
	Timestamp        int64  `xorm:"timestamp int notnull index"`
	WitnessSignature string `xorm:"witness_signature varchar(1024) notnull"`
	WitnessAddress   string `xorm:"witness_address char(64)"`
	TransactionRoot  string `xorm:"transaction_root char(64) notnull"`
	Size             int64  `xorm:"size int notnull"`
}

func (t Block) TableName() string {
	return tableName("block")
}
