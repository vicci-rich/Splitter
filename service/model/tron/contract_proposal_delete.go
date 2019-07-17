package tron

type ProposalDeleteContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	ProposalID      int64  `xorm:"proposal_id int"`
}

func (c ProposalDeleteContract) TableName() string {
	return tableName("contract_proposal_delete")
}
