package tron

type AssetIssueContract struct {
	ID                      int64  `xorm:"id bigint autoincr pk"`
	TransactionHash         string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber             int64  `xorm:"block_number int index"`
	Timestamp               int64  `xorm:"timestamp int notnull index"`
	AssetID                 string `xorm:"asset_id varchar(64)"`
	OwnerAddress            string `xorm:"owner_address char(64) notnull"`
	Name                    string `xorm:"name varchar(1024) notnull"`
	Abbr                    string `xorm:"abbr varchar(1024) notnull"`
	Description             string `xorm:"description varchar(1024) notnull"`
	Url                     string `xorm:"url varchar(1024) notnull"`
	FrozenAmount            int64  `xorm:"frozen_amount bigint notnull"`
	FrozenDays              int64  `xorm:"frozen_days int notnull"`
	TotalSupply             int64  `xorm:"total_supply bigint notnull"`
	StartTime               int64  `xorm:"start_time int notnull"`
	EndTime                 int64  `xorm:"end_time bigint notnull"`
	TrxNum                  int64  `xorm:"trx_num bigint"`
	Num                     int64  `xorm:"num bigint"`
	Precision               int64  `xorm:"precision int"`
	VoteScore               int64  `xorm:"vote_score int"`
	Order                   int64  `xorm:"order int"`
	FreeAssetNetLimit       int64  `xorm:"free_asset_net_limit int"`
	PublicFreeAssetNetLimit int64  `xorm:"public_free_asset_net_limit int"`
	PublicFreeAssetNetUsage int64  `xorm:"public_free_asset_net_usage int"`
	PublicLatestFreeNetTime int64  `xorm:"public_latest_free_net_time int"`
	FrozenSupplyNum         int    `xorm:"frozen_supply_num int notnull"`

	//repeated FrozenSupply frozen_supply = 5;
}

func (c AssetIssueContract) TableName() string {
	return tableName("contract_asset_issue")
}
