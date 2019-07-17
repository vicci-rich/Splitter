package bsv

type Block struct {
	ID              int64   `xorm:"id bigint autoincr pk"`
	Height          int64   `xorm:"height int notnull unique"`
	Size            int64   `xorm:"size int notnull"`
	Timestamp       int64   `xorm:"timestamp int notnull index"`
	Version         int64   `xorm:"version int notnull"`
	MerkleRoot      string  `xorm:"merkle_root char(64) notnull"`
	Bits            string  `xorm:"bits char(8) notnull"`
	Nonce           int64   `xorm:"nonce bigint notnull"`
	Hash            string  `xorm:"hash char(64) notnull index"`
	MedianTimestamp int64   `xorm:"median_timestamp int notnull"`
	Difficulty      float64 `xorm:"difficulty double notnull"`
	PreviousHash    string  `xorm:"prev_hash char(64) notnull"`
	ChainWork       string  `xorm:"chain_work char(64) notnull"`
	TxCount         int64   `xorm:"tx_count int notnull"`
	PoolName        string  `xorm:"pool_name varchar(100) notnull default ''"`
}

func (t Block) TableName() string {
	return tableName("block")
}
