package eos

type Block struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	BlockNum          int64  `xorm:"block_num int notnull unique index"`
	Hash              string `xorm:"hash char(64) notnull index"`
	Previous          string `xorm:"previous char(64) notnull"`
	Timestamp         int64  `xorm:"timestamp int notnull index"`
	TimestampISO      string `xorm:"timestampISO char(24) notnull"`
	ProducerSignature string `xorm:"producer_signature varchar(128) notnull"`
	Producer          string `xorm:"producer char(12)"`
	NewProducers      string `xorm:"new_producers text"`
	TransactionMRoot  string `xorm:"transaction_mroot char(64) notnull"`
	ActionMRoot       string `xorm:"action_mroot char(64) notnull"`
	ScheduleVersion   int64  `xorm:"schedule_version int notnull"`
	RefBlockPrefix    int64  `xorm:"ref_block_prefix bigint notnull"`
	TransactionLen    int64  `xorm:"tx_len int notnull"`
	Confirmed         int64  `xorm:"confirmed int notnull"`
}

func (t Block) TableName() string {
	return tableName("block")
}
