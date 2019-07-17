package tron

type ProposalCreateContract struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	TransactionHash string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber     int64  `xorm:"block_number int index"`
	Timestamp       int64  `xorm:"timestamp int notnull index"`
	OwnerAddress    string `xorm:"owner_address char(64) notnull"`
	Parameters      string `xorm:"parameters text"`
}

func (c ProposalCreateContract) TableName() string {
	return tableName("contract_proposal_create")
}
