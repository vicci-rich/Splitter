package btc

type StatisticsWeek struct {
	ID                        int64   `xorm:"id bigint autoincr pk"`
	Timestamp                 int64   `xorm:"timestamp int notnull unique"`
	UserActive                int64   `xorm:"user_active bigint notnull default '0'"`
	Retention                 float64 `xorm:"retention double notnull default '0'"`
	RatioOfMarketValue        float64 `xorm:"ratio_of_market_value double notnull default '0'"`
	ValueConsume              float64 `xorm:"value_consume double notnull default '0'"`
	TxValueSum                float64 `xorm:"tx_value_sum double notnull default '0'"`
	DifficultySum             float64 `xorm:"difficulty_sum double notnull default '0'"`
	UserCount                 int64   `xorm:"user_count bigint notnull default '0'"`
	UserNew                   int64   `xorm:"user_new bigint notnull default '0'"`
	UserNewActiveRate         float64 `xorm:"user_new_active_rate double notnull default '0'"`
	ActiveAddressCount        int64   `xorm:"active_address_count bigint notnull default '0'"`
	TxCount                   int64   `xorm:"tx_count bigint notnull default '0'"`
	MultisigCount             int64   `xorm:"multisig_count bigint notnull default '0'"`
	NonstandardCount          int64   `xorm:"nonstandard_count bigint notnull default '0'"`
	NulldataCount             int64   `xorm:"nulldata_count bigint notnull default '0'"`
	PubkeyCount               int64   `xorm:"pubkey_count bigint notnull default '0'"`
	PubkeyhashCount           int64   `xorm:"pubkeyhash_count bigint notnull default '0'"`
	ScripthashCount           int64   `xorm:"scripthash_count bigint notnull default '0'"`
	WithnessV0KeyhashCount    int64   `xorm:"withness_v0_keyhash_count bigint notnull default '0'"`
	WithnessV0ScripthashCount int64   `xorm:"withness_v0_scripthash_count bigint notnull default '0'"`
	ValueLevel1               int64   `xorm:"value_level1 bigint notnull default '0'"`
	ValueLevel2               int64   `xorm:"value_level2 bigint notnull default '0'"`
	ValueLevel3               int64   `xorm:"value_level3 bigint notnull default '0'"`
	ValueLevel4               int64   `xorm:"value_level4 bigint notnull default '0'"`
	ValueLevel5               int64   `xorm:"value_level5 bigint notnull default '0'"`
	ValueLevel6               int64   `xorm:"value_level6 bigint notnull default '0'"`
	ValueCountLevel1          int64   `xorm:"value_count_level1 bigint notnull default '0'"`
	ValueCountLevel2          int64   `xorm:"value_count_level2 bigint notnull default '0'"`
	ValueCountLevel3          int64   `xorm:"value_count_level3 bigint notnull default '0'"`
	ValueCountLevel4          int64   `xorm:"value_count_level4 bigint notnull default '0'"`
	ValueCountLevel5          int64   `xorm:"value_count_level5 bigint notnull default '0'"`
	ValueCountLevel6          int64   `xorm:"value_count_level6 bigint notnull default '0'"`
}

func (t StatisticsWeek) TableName() string {
	return tableName("statistics_week")
}
