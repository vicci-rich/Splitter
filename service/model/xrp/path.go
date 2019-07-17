package xrp

type Path struct {
	ID          int64  `xorm:"id bigint autoincr pk"`
	ParentHash  string `xorm:"parent_hash char(64) notnull index"`
	LedgerIndex int64  `xorm:"ledger_index int index"`
	CloseTime   int64  `xorm:"close_time int notnull"`
	Currency    string `xorm:"currency char(8) notnull index"`
	Issuer      string `xorm:"issuer char(34) notnull index"`
	Type        int64  `xorm:"type int null"`
	Account     string `xorm:"account char(34) null index"`
	InTxIndex   int64  `xorm:"in_tx_index int"`
	InPathIndex int64  `xorm:"in_path_index int"`
}

func (t Path) TableName() string {
	return tableName("path")
}
