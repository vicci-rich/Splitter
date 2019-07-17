package tron

type ExchangeCreateContract struct {
	ID                 int64  `xorm:"id bigint autoincr pk"`
	TransactionHash    string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber        int64  `xorm:"block_number int index"`
	Timestamp          int64  `xorm:"timestamp int notnull index"`
	OwnerAddress       string `xorm:"owner_address char(64) notnull"`
	FirstTokenID       string `xorm:"first_token_id varchar(64)"`
	FirstTokenBalance  int64  `xorm:"first_token_balance bigint"`
	SecondTokenID      string `xorm:"second_token_id varchar(64)"`
	SecondTokenBalance int64  `xorm:"second_token_balance bigint"`
}

func (c ExchangeCreateContract) TableName() string {
	return tableName("contract_exchange_create")
}
