package doge

import "fmt"

type VIn struct {
	ID               int64  `xorm:"id bigint autoincr pk"`
	BlockHeight      int64  `xorm:"block_height int notnull index(block_height)"`
	TxID             string `xorm:"tx_id char(64) notnull unique(tx_id_number)"`
	TxIDOrigin       string `xorm:"tx_id_origin char(64) notnull"`
	Timestamp        int64  `xorm:"timestamp int notnull index(timestamp)"`
	Coinbase         string `xorm:"coinbase varchar(200) notnull"`
	Sequence         int64  `xorm:"sequence bigint notnull"`
	Number           int64  `xorm:"number int notnull unique(tx_id_number)"`
	VOutNumberOrigin int64  `xorm:"vout_num_origin smallint notnull"`
	AddressOrigin    string `xorm:"address_origin varchar(256) notnull index(address_origin)"`
	ValueOrigin      int64  `xorm:"value_origin bigint notnull"`
	ScriptSignature  string `xorm:"script_sig text notnull"`
	TxInWitness      string `xorm:"tx_in_witness text notnull"`
}

func (t VIn) TableName() string {
	return tableName("vin")
}

func (t VIn) String() string {
	msg := fmt.Sprintf("id: %d\nheight: %d\ntxid: %s\ntxid_org:%s\nts: %d\ncoinbase: %s\nseq: %d\nnum: %d\n"+
		"vout_num_org: %d\naddr_org: %s\nval_org: %d\nscp_sig: %s\ntxin_wit: %s\n",
		t.ID, t.BlockHeight, t.TxID, t.TxIDOrigin, t.Timestamp, t.Coinbase, t.Sequence, t.Number, t.VOutNumberOrigin,
		t.AddressOrigin, t.ValueOrigin, t.ScriptSignature, t.TxIDOrigin,
	)
	return msg
}
