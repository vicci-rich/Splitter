package xlm

type Ledger struct {
	ID                         int64  `xorm:"id bigint autoincr pk"`
	LedgerID                   string `xorm:"ledger_id char(64) notnull index(IDX_xlm_ledger_id)"`
	PagingToken                string `xorm:"paging_token char(64) notnull"`
	LedgerHash                 string `xorm:"ledger_hash char(64) notnull index(IDX_xlm_ledger_hash)"`
	PreviousLedgerHash         string `xorm:"previous_ledger_hash char(64) notnull"`
	Sequence                   int64  `xorm:"sequence int notnull unique(IDX_xlm_ledger_sequence)"`
	TransactionCount           int64  `xorm:"transaction_count int notnull"`
	SuccessfulTransactionCount int64  `xorm:"successful_transaction_count int notnull"`
	FailedTransactionCount     int64  `xorm:"failed_transaction_count int notnull"`
	OperationCount             int64  `xorm:"operation_count int notnull"`
	ClosedTime                 int64  `xorm:"closed_time int notnull"`
	TotalCoins                 string `xorm:"total_coins char(64) notnull"`
	FeePool                    string `xorm:"fee_pool char(64) notnull"`
	BaseFeeInStroops           int64  `xrom:"base_fee_in_stroops int notnull"`
	BaseReserveInStroops       int64  `xrom:"base_reserve_in_stroops int notnull"`
	MaxTxSetSize               int64  `xorm:"max_tx_set_size int notnull"`
	ProtocolVersion            int64  `xorm:"protocol_version int notnull"`
}

func (t Ledger) TableName() string {
	return tableName("ledger")
}
