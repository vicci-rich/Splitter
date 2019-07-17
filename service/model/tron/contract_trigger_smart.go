package tron

type TriggerSmartContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	ContractAddress string `xorm:"contract_address char(64) notnull"`
	//Data            string `xorm:"data varchar(1024) notnull"`
	CallValue      int64 `xorm:"call_value bigint"`
	CallTokenValue int64 `xorm:"call_token_value bigint"`
	TokenID        int64 `xorm:"token_id int"`
}

func (c TriggerSmartContract) TableName() string {
	return tableName("contract_trigger_smart")
}
