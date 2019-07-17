package tron

type ParticipateAssetIssueContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	ToAddress       string `xorm:"to_address char(64) notnull"`
	AssetName       string `xorm:"asset_name varchar(1024)"`
	Amount          int64  `xorm:"amount bigint"`
}

func (c ParticipateAssetIssueContract) TableName() string {
	return tableName("contract_participate_asset_issue")
}
