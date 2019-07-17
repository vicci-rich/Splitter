package tron

type TransferAssetContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	ToAddress       string `xorm:"to_address char(64) notnull"`
	AssetName       string `xorm:"asset_name varchar(1024) notnull"`
	Amount          int64  `xorm:"amount bigint notnull"`
}

func (c TransferAssetContract) TableName() string {
	return tableName("contract_transfer_asset")
}
