package xrp

type Amount struct {
	ID          int64  `xorm:"id bigint autoincr pk"`
	ParentHash  string `xorm:"parent_hash char(64) notnull index"`
	LedgerIndex int64  `xorm:"ledger_index int index"`
	CloseTime   int64  `xorm:"close_time int notnull"`
	Currency    string `xorm:"currency char(8) notnull index"`
	Value       string `xorm:"value decimal(38,4) notnull default '0'"`
	Issuer      string `xorm:"issuer char(34) notnull index"`
	AmountType  int    `xorm:"amount_type int notnull"`
	//1:Payment Amount  2:Payment SendMax 3:Payment DeliverMin
	//4: OfferCreate TakerGets  5:OfferCreate TakerPay
	//6: TrustSet LimitAmount

}

func (t Amount) TableName() string {
	return tableName("amount")
}
