package xlm

type Operation struct {
	ID               int64  `xorm:"id bigint autoincr pk"`
	OperationID      string `xorm:"operation_id char(64) notnull index(IDX_xlm_operation_id)"`
	TransactionID    string `xorm:"transaction_id char(64) notnull index(IDX_xlm_transaction_id)"`
	ApplicationOrder int64  `xorm:"application_order int notnull"`
	Type             string `xorm:"type char(32) notnull"`
	Detail           string `xorm:"detail text notnull"`
	SourceAccount    string `xorm:"source_account char(64) notnull index(IDX_xlm_source_account_hash)"`
}

func (t Operation) TableName() string {
	return tableName("operation")
}
