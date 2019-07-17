package tron

type Transaction struct {
	ID          int64       `xorm:"id bigint autoincr pk"`
	Hash        string      `xorm:"hash char(64) notnull index"`
	BlockNumber int64       `xorm:"block_number int index"`
	Timestamp   int64       `xorm:"timestamp int notnull index"`
	Contracts   []*Contract `xorm:"-"`
}

func (t Transaction) TableName() string {
	return tableName("transaction")
}

type Contract struct {
	BlockNumber     int64
	TransactionHash string
	ContractNumber  int
	Type            string
	Value           string
}
