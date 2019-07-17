package tron

type ExchangeTransactionContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	TokenID         string `xorm:"token_id varchar(512)"`
	ExchangeID      int64  `xorm:"exchange_id int"`
	Quant           int64  `xorm:"quant bigint"`
	Expected        int64  `xorm:"expected bigint"`
}

func (c ExchangeTransactionContract) TableName() string {
	return tableName("contract_exchange_transaction")
}
