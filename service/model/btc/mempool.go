package btc

type Mempool struct {
	ID        int64   `xorm:"id bigint autoincr pk"`
	Timestamp int64   `xorm:"timestamp bigint notnull index"`
	Count     int64   `xorm:"count bigint notnull"`
	Bytes     int64   `xorm:"bytes bigint notnull"`
	Rate      float64 `xorm:"rate double notnull"`
	BestFee   float64 `xorm:"best_fee double notnull"`
}

func (t Mempool) TableName() string {
	return tableName("mempool")
}
