package tron

type UpdateEnergyLimitContract struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	TransactionHash   string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber       int64  `xorm:"block_number int index"`
	Timestamp         int64  `xorm:"timestamp int notnull index"`
	OwnerAddress      string `xorm:"owner_address char(64) notnull"`
	ContractAddress   string `xorm:"contract_address char(64) notnull"`
	OriginEnergyLimit int64  `xorm:"origin_energy_limit bigint notnull"`
}

func (c UpdateEnergyLimitContract) TableName() string {
	return tableName("contract_update_energy_limit")
}
