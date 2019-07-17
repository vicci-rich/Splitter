package xrp

type AffectedNodes struct {
	ID          int64  `xorm:"id bigint autoincr pk"`
	ParentHash  string `xorm:"parent_hash char(64) notnull index"`
	LedgerIndex int64  `xorm:"ledger_index int index"`
	CloseTime   int64  `xorm:"close_time int notnull"`

	NodeType          string `xorm:"node_type char(20) notnull index"`
	LedgerEntryType   string `xorm:"ledger_entry_type char(20) null index"`
	NodeLedgerIndex   string `xorm:"hash char(64) not null"`
	PreviousTxnID     string `xorm:"previous_txn_id char(64) null"`
	PreviousTxnLgrSeq int64  `xorm:"previous_txn_lgr_seq int null"`
	FullJsonStr       string `xorm:"full_json_str varchar(1024) null"`
}

func (t AffectedNodes) TableName() string {
	return tableName("transaction_affected_nodes")
}
