package tron

type CreateSmartContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	ContractAddress string `xorm:"contract_address char(64) notnull"`
	Name            string `xorm:"name varchar(1024)"`
	TokenID         int64  `xorm:"token_id int"`
	CallTokenValue  int64  `xorm:"call_token_value bigint"`
}

func (c CreateSmartContract) TableName() string {
	return tableName("contract_create_smart")
}
