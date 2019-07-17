package xlm

type Transaction struct {
	ID                    int64  `xorm:"id bigint autoincr pk"`
	TransactionID         string `xorm:"transaction_id char(64) notnull index(IDX_xlm_transaction_id)"`
	PagingToken           string `xorm:"paging_token char(64) notnull"`
	TransactionHash       string `xorm:"transaction_hash char(64) notnull index(IDX_xlm_transaction_hash)"`
	LedgerSequence        int64  `xorm:"ledger_sequence int notnull index(IDX_xlm_transaction_ledger_sequence)"`
	SourceAccount         string `xorm:"source_account char(64) notnull index(IDX_xlm_source_account_hash)"`
	SourceAccountSequence string `xorm:"source_account_sequence char(64) notnull index(IDX_xlm_source_account_sequence_hash)"`
	FeePaid               int64  `xorm:"fee_paid int notnull"`
	OperationCount        int64  `xorm:"operation_count int notnull"`
	MemoType              string `xorm:"memo_type char(64) notnull"`
	Signatures            string `xorm:"signatures text notnull"`
}

func (t Transaction) TableName() string {
	return tableName("transaction")
}
