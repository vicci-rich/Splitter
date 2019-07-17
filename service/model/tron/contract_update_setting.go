package tron

type UpdateSettingContract struct {
	ID                         int64  `xorm:"id bigint autoincr pk"`
	TransactionHash            string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber                int64  `xorm:"block_number int index"`
	Timestamp                  int64  `xorm:"timestamp int notnull index"`
	OwnerAddress               string `xorm:"owner_address char(64) notnull"`
	ContractAddress            string `xorm:"contract_address char(64) notnull"`
	ConsumeUserResourcePercent int64  `xorm:"consume_user_resource_percent int"`
}

func (c UpdateSettingContract) TableName() string {
	return tableName("contract_update_setting")
}
