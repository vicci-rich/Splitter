package core

type CurrencyExchangeRate struct {
	ID        int64   `xorm:"id bigint autoincr pk"`
	Timestamp int     `xorm:"timestamp int notnull unique"`
	USD       float64 `xorm:"usd double notnull"`
	CNY       float64 `xorm:"cny double notnull"`
	EUR       float64 `xorm:"eur double notnull"`
	JPY       float64 `xorm:"jpy double notnull"`
	KRW       float64 `xorm:"krw double notnull"`
}

func (t CurrencyExchangeRate) TableName() string {
	return "currency_exchange_rate"
}
