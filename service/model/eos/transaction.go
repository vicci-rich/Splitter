package eos

type Transaction struct {
	ID                    int64  `xorm:"id bigint autoincr pk"`
	Hash                  string `xorm:"hash char(64) notnull index"`
	BlockNum              int64  `xorm:"block_num int index"`
	Status                string `xorm:"status varchar(24) notnull"`
	Compression           string `xorm:"compression varchar(40) notnull"`
	Expiration            string `xorm:"expiration char(24) notnull"`
	CpuUsageUs            int64  `xorm:"cpu_usage_us int notnull"`
	NetUsageWords         int64  `xorm:"net_usage_words int notnull"`
	DelaySec              int64  `xorm:"delay_sec int notnull"`
	MaxCpuUsageMs         int64  `xorm:"max_cpu_usage_ms int notnull"`
	MaxNetUsageWords      int64  `xorm:"max_net_usage_words int notnull"`
	PackedContextFreeData string `xorm:"packed_context_free_data varchar(512)"`
	PackedTrx             string `xorm:"packed_trx text"`
	RefBlockNum           int64  `xorm:"ref_block_num int notnull"`
	RefBlockPrefix        int64  `xorm:"ref_block_prefix bigint notnull"`
}

func (t Transaction) TableName() string {
	return tableName("transaction")
}
