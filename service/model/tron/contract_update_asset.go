package tron

type UpdateAssetContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	Description     string `xorm:"description varchar(4096)"`
	Url             string `xorm:"url varchar(1024)"`
	NewLimit        int64  `xorm:"new_limit bigint"`
	NewPublicLimit  int64  `xorm:"new_public_limit bigint"`
}

func (c UpdateAssetContract) TableName() string {
	return tableName("contract_update_asset")
}
