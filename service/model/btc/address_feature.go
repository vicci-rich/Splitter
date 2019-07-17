package btc

type AddressFeature struct {
	ID              int64   `xorm:"id bigint autoincr pk"`
	Address         string  `xorm:"address varchar(255) notnull unique"`
	InputTxNumber   int64   `xorm:"input_tx_number int notnull default '0'"`
	OutputTxNumber  int64   `xorm:"output_tx_number int notnull default '0'"`
	InputValue      float64 `xorm:"input_value double notnull default '0'"`
	OutputValue     float64 `xorm:"output_value double notnull default '0'"`
	InputNumberAvg  float64 `xorm:"input_number_avg double notnull default '0'"`
	OutputNumberAvg float64 `xorm:"output_number_avg double notnull default '0'"`
	ProportionA     float64 `xorm:"proportion_a double notnull default '0'"`
	ProportionB     float64 `xorm:"proportion_b double notnull default '0'"`
	InValueAvg      float64 `xorm:"in_value_avg double notnull default '0'"`
	OutValueAvg     float64 `xorm:"out_value_avg double notnull default '0'"`
	IsMining        int     `xorm:"is_mining int notnull default '0'"`
	InFeeTotal      float64 `xorm:"in_fee_total double notnull default '0'"`
	OutFeeTotal     float64 `xorm:"out_fee_total double notnull default '0'"`
	InFeeAvg        float64 `xorm:"in_fee_avg double notnull default '0'"`
	OutFeeAvg       float64 `xorm:"out_fee_avg double notnull default '0'"`
	InTxDayAvg      float64 `xorm:"in_tx_day_avg double notnull default '0'"`
	OutTxDayAvg     float64 `xorm:"out_tx_day_avg double notnull default '0'"`
	Category        string  `xorm:"category varchar(20) notnull"`
}

func (t AddressFeature) TableName() string {
	return tableName("address_feature")
}
