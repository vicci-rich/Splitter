package eth

type Mempool struct {
	ID           int64 `xorm:"id bigint autoincr pk"`
	Timestamp    int64 `xorm:"timestamp bigint notnull index"`
	PendingCount int64 `xorm:"count bigint notnull"`
}

func (t Mempool) TableName() string {
	return tableName("mempool")
}
