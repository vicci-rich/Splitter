package tron

type FreezeBalanceContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	FrozenBalance   int64  `xorm:"frozen_balance bigint notnull"`
	FrozenDuration  int64  `xorm:"frozen_duration int notnull"`
}

func (c FreezeBalanceContract) TableName() string {
	return tableName("contract_freeze_balance")
}
