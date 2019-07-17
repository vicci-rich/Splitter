package eth

import (
	"math/big"
)

type StatisticsDayAdditionalTokenIndexes struct {
	ID        int64 `xorm:"id bigint autoincr pk"`
	Timestamp int64 `xorm:"timestamp int notnull unique index"`

	BNBLargeTransactionCount  int64 `xorm:"bnb_large_transaction_count int null default '0'"`
	MKRLargeTransactionCount  int64 `xorm:"mkr_large_transaction_count int null default '0'"`
	USDCLargeTransactionCount int64 `xorm:"usdc_large_transaction_count int null default '0'"`
	TUSDLargeTransactionCount int64 `xorm:"tusd_large_transaction_count int null default '0'"`
	GUSDLargeTransactionCount int64 `xorm:"gusd_large_transaction_count int null default '0'"`
	HTLargeTransactionCount   int64 `xorm:"ht_large_transaction_count int null default '0'"`
	OMGLargeTransactionCount  int64 `xorm:"omg_large_transaction_count int null default '0'"`
	ZRXLargeTransactionCount  int64 `xorm:"zrx_large_transaction_count int null default '0'"`
	BATLargeTransactionCount  int64 `xorm:"bat_large_transaction_count int null default '0'"`

	BNBTransactionValueInterval1  int64 `xorm:"bnb_transaction_value_interval_1 int null default '0'"`
	MKRTransactionValueInterval1  int64 `xorm:"mkr_transaction_value_interval_1 int null default '0'"`
	USDCTransactionValueInterval1 int64 `xorm:"usdc_transaction_value_interval_1 int null default '0'"`
	TUSDTransactionValueInterval1 int64 `xorm:"tusd_transaction_value_interval_1 int null default '0'"`
	GUSDTransactionValueInterval1 int64 `xorm:"gusd_transaction_value_interval_1 int null default '0'"`
	HTTransactionValueInterval1   int64 `xorm:"ht_transaction_value_interval_1 int null default '0'"`
	OMGTransactionValueInterval1  int64 `xorm:"omg_transaction_value_interval_1 int null default '0'"`
	ZRXTransactionValueInterval1  int64 `xorm:"zrx_transaction_value_interval_1 int null default '0'"`
	BATTransactionValueInterval1  int64 `xorm:"bat_transaction_value_interval_1 int null default '0'"`

	BNBTransactionValueInterval2  int64 `xorm:"bnb_transaction_value_interval_2 int null default '0'"`
	MKRTransactionValueInterval2  int64 `xorm:"mkr_transaction_value_interval_2 int null default '0'"`
	USDCTransactionValueInterval2 int64 `xorm:"usdc_transaction_value_interval_2 int null default '0'"`
	TUSDTransactionValueInterval2 int64 `xorm:"tusd_transaction_value_interval_2 int null default '0'"`
	GUSDTransactionValueInterval2 int64 `xorm:"gusd_transaction_value_interval_2 int null default '0'"`
	HTTransactionValueInterval2   int64 `xorm:"ht_transaction_value_interval_2 int null default '0'"`
	OMGTransactionValueInterval2  int64 `xorm:"omg_transaction_value_interval_2 int null default '0'"`
	ZRXTransactionValueInterval2  int64 `xorm:"zrx_transaction_value_interval_2 int null default '0'"`
	BATTransactionValueInterval2  int64 `xorm:"bat_transaction_value_interval_2 int null default '0'"`

	BNBTransactionValueInterval3  int64 `xorm:"bnb_transaction_value_interval_3 int null default '0'"`
	MKRTransactionValueInterval3  int64 `xorm:"mkr_transaction_value_interval_3 int null default '0'"`
	USDCTransactionValueInterval3 int64 `xorm:"usdc_transaction_value_interval_3 int null default '0'"`
	TUSDTransactionValueInterval3 int64 `xorm:"tusd_transaction_value_interval_3 int null default '0'"`
	GUSDTransactionValueInterval3 int64 `xorm:"gusd_transaction_value_interval_3 int null default '0'"`
	HTTransactionValueInterval3   int64 `xorm:"ht_transaction_value_interval_3 int null default '0'"`
	OMGTransactionValueInterval3  int64 `xorm:"omg_transaction_value_interval_3 int null default '0'"`
	ZRXTransactionValueInterval3  int64 `xorm:"zrx_transaction_value_interval_3 int null default '0'"`
	BATTransactionValueInterval3  int64 `xorm:"bat_transaction_value_interval_3 int null default '0'"`

	BNBTransactionValueInterval4  int64 `xorm:"bnb_transaction_value_interval_4 int null default '0'"`
	MKRTransactionValueInterval4  int64 `xorm:"mkr_transaction_value_interval_4 int null default '0'"`
	USDCTransactionValueInterval4 int64 `xorm:"usdc_transaction_value_interval_4 int null default '0'"`
	TUSDTransactionValueInterval4 int64 `xorm:"tusd_transaction_value_interval_4 int null default '0'"`
	GUSDTransactionValueInterval4 int64 `xorm:"gusd_transaction_value_interval_4 int null default '0'"`
	HTTransactionValueInterval4   int64 `xorm:"ht_transaction_value_interval_4 int null default '0'"`
	OMGTransactionValueInterval4  int64 `xorm:"omg_transaction_value_interval_4 int null default '0'"`
	ZRXTransactionValueInterval4  int64 `xorm:"zrx_transaction_value_interval_4 int null default '0'"`
	BATTransactionValueInterval4  int64 `xorm:"bat_transaction_value_interval_4 int null default '0'"`

	BNBTransactionValueInterval5  int64 `xorm:"bnb_transaction_value_interval_5 int null default '0'"`
	MKRTransactionValueInterval5  int64 `xorm:"mkr_transaction_value_interval_5 int null default '0'"`
	USDCTransactionValueInterval5 int64 `xorm:"usdc_transaction_value_interval_5 int null default '0'"`
	TUSDTransactionValueInterval5 int64 `xorm:"tusd_transaction_value_interval_5 int null default '0'"`
	GUSDTransactionValueInterval5 int64 `xorm:"gusd_transaction_value_interval_5 int null default '0'"`
	HTTransactionValueInterval5   int64 `xorm:"ht_transaction_value_interval_5 int null default '0'"`
	OMGTransactionValueInterval5  int64 `xorm:"omg_transaction_value_interval_5 int null default '0'"`
	ZRXTransactionValueInterval5  int64 `xorm:"zrx_transaction_value_interval_5 int null default '0'"`
	BATTransactionValueInterval5  int64 `xorm:"bat_transaction_value_interval_5 int null default '0'"`

	BNBTransactionValueInterval6  int64 `xorm:"bnb_transaction_value_interval_6 int null default '0'"`
	MKRTransactionValueInterval6  int64 `xorm:"mkr_transaction_value_interval_6 int null default '0'"`
	USDCTransactionValueInterval6 int64 `xorm:"usdc_transaction_value_interval_6 int null default '0'"`
	TUSDTransactionValueInterval6 int64 `xorm:"tusd_transaction_value_interval_6 int null default '0'"`
	GUSDTransactionValueInterval6 int64 `xorm:"gusd_transaction_value_interval_6 int null default '0'"`
	HTTransactionValueInterval6   int64 `xorm:"ht_transaction_value_interval_6 int null default '0'"`
	OMGTransactionValueInterval6  int64 `xorm:"omg_transaction_value_interval_6 int null default '0'"`
	ZRXTransactionValueInterval6  int64 `xorm:"zrx_transaction_value_interval_6 int null default '0'"`
	BATTransactionValueInterval6  int64 `xorm:"bat_transaction_value_interval_6 int null default '0'"`

	BNBAccountBalanceInterval1  int64 `xorm:"bnb_account_balance_interval_1 int null default '0'"`
	MKRAccountBalanceInterval1  int64 `xorm:"mkr_account_balance_interval_1 int null default '0'"`
	USDCAccountBalanceInterval1 int64 `xorm:"usdc_account_balance_interval_1 int null default '0'"`
	TUSDAccountBalanceInterval1 int64 `xorm:"tusd_account_balance_interval_1 int null default '0'"`
	GUSDAccountBalanceInterval1 int64 `xorm:"gusd_account_balance_interval_1 int null default '0'"`
	HTAccountBalanceInterval1   int64 `xorm:"ht_account_balance_interval_1 int null default '0'"`
	OMGAccountBalanceInterval1  int64 `xorm:"omg_account_balance_interval_1 int null default '0'"`
	ZRXAccountBalanceInterval1  int64 `xorm:"zrx_account_balance_interval_1 int null default '0'"`
	BATAccountBalanceInterval1  int64 `xorm:"bat_account_balance_interval_1 int null default '0'"`

	BNBAccountBalanceInterval2  int64 `xorm:"bnb_account_balance_interval_2 int null default '0'"`
	MKRAccountBalanceInterval2  int64 `xorm:"mkr_account_balance_interval_2 int null default '0'"`
	USDCAccountBalanceInterval2 int64 `xorm:"usdc_account_balance_interval_2 int null default '0'"`
	TUSDAccountBalanceInterval2 int64 `xorm:"tusd_account_balance_interval_2 int null default '0'"`
	GUSDAccountBalanceInterval2 int64 `xorm:"gusd_account_balance_interval_2 int null default '0'"`
	HTAccountBalanceInterval2   int64 `xorm:"ht_account_balance_interval_2 int null default '0'"`
	OMGAccountBalanceInterval2  int64 `xorm:"omg_account_balance_interval_2 int null default '0'"`
	ZRXAccountBalanceInterval2  int64 `xorm:"zrx_account_balance_interval_2 int null default '0'"`
	BATAccountBalanceInterval2  int64 `xorm:"bat_account_balance_interval_2 int null default '0'"`

	BNBAccountBalanceInterval3  int64 `xorm:"bnb_account_balance_interval_3 int null default '0'"`
	MKRAccountBalanceInterval3  int64 `xorm:"mkr_account_balance_interval_3 int null default '0'"`
	USDCAccountBalanceInterval3 int64 `xorm:"usdc_account_balance_interval_3 int null default '0'"`
	TUSDAccountBalanceInterval3 int64 `xorm:"tusd_account_balance_interval_3 int null default '0'"`
	GUSDAccountBalanceInterval3 int64 `xorm:"gusd_account_balance_interval_3 int null default '0'"`
	HTAccountBalanceInterval3   int64 `xorm:"ht_account_balance_interval_3 int null default '0'"`
	OMGAccountBalanceInterval3  int64 `xorm:"omg_account_balance_interval_3 int null default '0'"`
	ZRXAccountBalanceInterval3  int64 `xorm:"zrx_account_balance_interval_3 int null default '0'"`
	BATAccountBalanceInterval3  int64 `xorm:"bat_account_balance_interval_3 int null default '0'"`

	BNBAccountBalanceInterval4  int64 `xorm:"bnb_account_balance_interval_4 int null default '0'"`
	MKRAccountBalanceInterval4  int64 `xorm:"mkr_account_balance_interval_4 int null default '0'"`
	USDCAccountBalanceInterval4 int64 `xorm:"usdc_account_balance_interval_4 int null default '0'"`
	TUSDAccountBalanceInterval4 int64 `xorm:"tusd_account_balance_interval_4 int null default '0'"`
	GUSDAccountBalanceInterval4 int64 `xorm:"gusd_account_balance_interval_4 int null default '0'"`
	HTAccountBalanceInterval4   int64 `xorm:"ht_account_balance_interval_4 int null default '0'"`
	OMGAccountBalanceInterval4  int64 `xorm:"omg_account_balance_interval_4 int null default '0'"`
	ZRXAccountBalanceInterval4  int64 `xorm:"zrx_account_balance_interval_4 int null default '0'"`
	BATAccountBalanceInterval4  int64 `xorm:"bat_account_balance_interval_4 int null default '0'"`

	BNBAccountBalanceInterval5  int64 `xorm:"bnb_account_balance_interval_5 int null default '0'"`
	MKRAccountBalanceInterval5  int64 `xorm:"mkr_account_balance_interval_5 int null default '0'"`
	USDCAccountBalanceInterval5 int64 `xorm:"usdc_account_balance_interval_5 int null default '0'"`
	TUSDAccountBalanceInterval5 int64 `xorm:"tusd_account_balance_interval_5 int null default '0'"`
	GUSDAccountBalanceInterval5 int64 `xorm:"gusd_account_balance_interval_5 int null default '0'"`
	HTAccountBalanceInterval5   int64 `xorm:"ht_account_balance_interval_5 int null default '0'"`
	OMGAccountBalanceInterval5  int64 `xorm:"omg_account_balance_interval_5 int null default '0'"`
	ZRXAccountBalanceInterval5  int64 `xorm:"zrx_account_balance_interval_5 int null default '0'"`
	BATAccountBalanceInterval5  int64 `xorm:"bat_account_balance_interval_5 int null default '0'"`

	BNBAccountBalanceInterval6  int64 `xorm:"bnb_account_balance_interval_6 int null default '0'"`
	MKRAccountBalanceInterval6  int64 `xorm:"mkr_account_balance_interval_6 int null default '0'"`
	USDCAccountBalanceInterval6 int64 `xorm:"usdc_account_balance_interval_6 int null default '0'"`
	TUSDAccountBalanceInterval6 int64 `xorm:"tusd_account_balance_interval_6 int null default '0'"`
	GUSDAccountBalanceInterval6 int64 `xorm:"gusd_account_balance_interval_6 int null default '0'"`
	HTAccountBalanceInterval6   int64 `xorm:"ht_account_balance_interval_6 int null default '0'"`
	OMGAccountBalanceInterval6  int64 `xorm:"omg_account_balance_interval_6 int null default '0'"`
	ZRXAccountBalanceInterval6  int64 `xorm:"zrx_account_balance_interval_6 int null default '0'"`
	BATAccountBalanceInterval6  int64 `xorm:"bat_account_balance_interval_6 int null default '0'"`

	BNBAccountBalanceInterval7  int64 `xorm:"bnb_account_balance_interval_7 int null default '0'"`
	MKRAccountBalanceInterval7  int64 `xorm:"mkr_account_balance_interval_7 int null default '0'"`
	USDCAccountBalanceInterval7 int64 `xorm:"usdc_account_balance_interval_7 int null default '0'"`
	TUSDAccountBalanceInterval7 int64 `xorm:"tusd_account_balance_interval_7 int null default '0'"`
	GUSDAccountBalanceInterval7 int64 `xorm:"gusd_account_balance_interval_7 int null default '0'"`
	HTAccountBalanceInterval7   int64 `xorm:"ht_account_balance_interval_7 int null default '0'"`
	OMGAccountBalanceInterval7  int64 `xorm:"omg_account_balance_interval_7 int null default '0'"`
	ZRXAccountBalanceInterval7  int64 `xorm:"zrx_account_balance_interval_7 int null default '0'"`
	BATAccountBalanceInterval7  int64 `xorm:"bat_account_balance_interval_7 int null default '0'"`

	BNBAccountBalanceInterval8  int64 `xorm:"bnb_account_balance_interval_8 int null default '0'"`
	MKRAccountBalanceInterval8  int64 `xorm:"mkr_account_balance_interval_8 int null default '0'"`
	USDCAccountBalanceInterval8 int64 `xorm:"usdc_account_balance_interval_8 int null default '0'"`
	TUSDAccountBalanceInterval8 int64 `xorm:"tusd_account_balance_interval_8 int null default '0'"`
	GUSDAccountBalanceInterval8 int64 `xorm:"gusd_account_balance_interval_8 int null default '0'"`
	HTAccountBalanceInterval8   int64 `xorm:"ht_account_balance_interval_8 int null default '0'"`
	OMGAccountBalanceInterval8  int64 `xorm:"omg_account_balance_interval_8 int null default '0'"`
	ZRXAccountBalanceInterval8  int64 `xorm:"zrx_account_balance_interval_8 int null default '0'"`
	BATAccountBalanceInterval8  int64 `xorm:"bat_account_balance_interval_8 int null default '0'"`

	BNBAccountBalanceInterval9  int64 `xorm:"bnb_account_balance_interval_9 int null default '0'"`
	MKRAccountBalanceInterval9  int64 `xorm:"mkr_account_balance_interval_9 int null default '0'"`
	USDCAccountBalanceInterval9 int64 `xorm:"usdc_account_balance_interval_9 int null default '0'"`
	TUSDAccountBalanceInterval9 int64 `xorm:"tusd_account_balance_interval_9 int null default '0'"`
	GUSDAccountBalanceInterval9 int64 `xorm:"gusd_account_balance_interval_9 int null default '0'"`
	HTAccountBalanceInterval9   int64 `xorm:"ht_account_balance_interval_9 int null default '0'"`
	OMGAccountBalanceInterval9  int64 `xorm:"omg_account_balance_interval_9 int null default '0'"`
	ZRXAccountBalanceInterval9  int64 `xorm:"zrx_account_balance_interval_9 int null default '0'"`
	BATAccountBalanceInterval9  int64 `xorm:"bat_account_balance_interval_9 int null default '0'"`

	BNBAccountBalanceInterval10  int64 `xorm:"bnb_account_balance_interval_10 int null default '0'"`
	MKRAccountBalanceInterval10  int64 `xorm:"mkr_account_balance_interval_10 int null default '0'"`
	USDCAccountBalanceInterval10 int64 `xorm:"usdc_account_balance_interval_10 int null default '0'"`
	TUSDAccountBalanceInterval10 int64 `xorm:"tusd_account_balance_interval_10 int null default '0'"`
	GUSDAccountBalanceInterval10 int64 `xorm:"gusd_account_balance_interval_10 int null default '0'"`
	HTAccountBalanceInterval10   int64 `xorm:"ht_account_balance_interval_10 int null default '0'"`
	OMGAccountBalanceInterval10  int64 `xorm:"omg_account_balance_interval_10 int null default '0'"`
	ZRXAccountBalanceInterval10  int64 `xorm:"zrx_account_balance_interval_10 int null default '0'"`
	BATAccountBalanceInterval10  int64 `xorm:"bat_account_balance_interval_10 int null default '0'"`

	BNBDeadAccountNumberByThen  int64 `xorm:"bnb_dead_account_number_by_then int null default '0'"`
	MKRDeadAccountNumberByThen  int64 `xorm:"mkr_dead_account_number_by_then int null default '0'"`
	USDCDeadAccountNumberByThen int64 `xorm:"usdc_dead_account_number_by_then int null default '0'"`
	TUSDDeadAccountNumberByThen int64 `xorm:"tusd_dead_account_number_by_then int null default '0'"`
	GUSDDeadAccountNumberByThen int64 `xorm:"gusd_dead_account_number_by_then int null default '0'"`
	HTDeadAccountNumberByThen   int64 `xorm:"ht_dead_account_number_by_then int null default '0'"`
	OMGDeadAccountNumberByThen  int64 `xorm:"omg_dead_account_number_by_then int null default '0'"`
	ZRXDeadAccountNumberByThen  int64 `xorm:"zrx_dead_account_number_by_then int null default '0'"`
	BATDeadAccountNumberByThen  int64 `xorm:"bat_dead_account_number_by_then int null default '0'"`

	ActiveAddress      int64 `xorm:"active_address decimal(38,0) null default '0'"`
	NewAddress0        int64 `xorm:"new_address_0 decimal(38,0) null default '0'"`
	NewAddress1To7     int64 `xorm:"new_address_1_7 decimal(38,0) null default '0'"`
	NewAddress8To14    int64 `xorm:"new_address_8_14 decimal(38,0) null default '0'"`
	NewAddress15To21   int64 `xorm:"new_address_15_21 decimal(38,0) null default '0'"`
	NewAddress22To28   int64 `xorm:"new_address_22_28 decimal(38,0) null default '0'"`
	NewAddress29ToPlus int64 `xorm:"new_address_29_plus decimal(38,0) null default '0'"`

	TransactionValueInterval1 *big.Int `xorm:"transaction_value_interval_1 decimal(38,0) null default '0'"`
	TransactionValueInterval2 *big.Int `xorm:"transaction_value_interval_2 decimal(38,0) null default '0'"`
	TransactionValueInterval3 *big.Int `xorm:"transaction_value_interval_3 decimal(38,0) null default '0'"`
	TransactionValueInterval4 *big.Int `xorm:"transaction_value_interval_4 decimal(38,0) null default '0'"`
	TransactionValueInterval5 *big.Int `xorm:"transaction_value_interval_5 decimal(38,0) null default '0'"`
	TransactionValueInterval6 *big.Int `xorm:"transaction_value_interval_6 decimal(38,0) null default '0'"`

	TransactionCountInterval1 int64 `xorm:"transaction_count_interval_1 int null default '0'"`
	TransactionCountInterval2 int64 `xorm:"transaction_count_interval_2 int null default '0'"`
	TransactionCountInterval3 int64 `xorm:"transaction_count_interval_3 int null default '0'"`
	TransactionCountInterval4 int64 `xorm:"transaction_count_interval_4 int null default '0'"`
	TransactionCountInterval5 int64 `xorm:"transaction_count_interval_5 int null default '0'"`
	TransactionCountInterval6 int64 `xorm:"transaction_count_interval_6 int null default '0'"`

	ActiveCountIn1Days  int64 `xorm:"active_count_in_1_days int null default '0'"`
	ActiveCountIn2Days  int64 `xorm:"active_count_in_2_days int null default '0'"`
	ActiveCountIn3Days  int64 `xorm:"active_count_in_3_days int null default '0'"`
	ActiveCountIn4Days  int64 `xorm:"active_count_in_4_days int null default '0'"`
	ActiveCountIn5Days  int64 `xorm:"active_count_in_5_days int null default '0'"`
	ActiveCountIn6Days  int64 `xorm:"active_count_in_6_days int null default '0'"`
	ActiveCountIn7Days  int64 `xorm:"active_count_in_7_days int null default '0'"`
	ActiveCountIn8Days  int64 `xorm:"active_count_in_8_days int null default '0'"`
	ActiveCountIn9Days  int64 `xorm:"active_count_in_9_days int null default '0'"`
	ActiveCountIn10Days int64 `xorm:"active_count_in_10_days int null default '0'"`
	ActiveCountIn11Days int64 `xorm:"active_count_in_11_days int null default '0'"`
	ActiveCountIn12Days int64 `xorm:"active_count_in_12_days int null default '0'"`
	ActiveCountIn13Days int64 `xorm:"active_count_in_13_days int null default '0'"`
	ActiveCountIn14Days int64 `xorm:"active_count_in_14_days int null default '0'"`
	ActiveCountIn15Days int64 `xorm:"active_count_in_15_days int null default '0'"`

	LargeTransactionCountPercentage float64 `xorm:"large_transaction_count_percentage decimal(38,4) null default '0'"`
	LargeTransactionValuePercentage float64 `xorm:"large_transaction_value_percentage decimal(38,4) null default '0'"`

	ReviveAccountNumber     int64 `xorm:"revive_account_number int null default '0'"`
	ReliveAccountNumber     int64 `xorm:"relive_account_number int null default '0'"`
	DeadAccountNumberByThen int64 `xorm:"dead_account_number_by_then int null default '0'"`

	ContractTransactionCount int64 `xorm:"contract_transaction_count int null default '0'"`
}

func (t StatisticsDayAdditionalTokenIndexes) TableName() string {
	return tableName("statistics_day_additional_token_indexes")
}
