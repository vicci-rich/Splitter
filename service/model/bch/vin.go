package bch

import "fmt"

/*
DROP TABLE IF EXISTS `btc_vin`;
CREATE TABLE `btc_vin` (
  `id`             INT(10) UNSIGNED     NOT NULL AUTO_INCREMENT,
  `b_txid`         CHAR(64)             NOT NULL COMMENT '交易id，固定256bit，返回值用64位字符串表示16进制的哈希值，BLOCK数据',
  `c_number`       SMALLINT(5) UNSIGNED NOT NULL COMMENT '交易在区块中的顺序，计算数据',
  `b_block_height` INT(10) UNSIGNED     NOT NULL COMMENT '区块高度，BLOCK数据',
  `b_coinbase`     VARCHAR(200)         NOT NULL COMMENT 'coinbase记录的信息,2-100字节，coinbase交易独有，BLOCK数据',
  `b_txid_org`     CHAR(64)             NOT NULL COMMENT '输入来源的交易id，固定256bit，返回值用64位字符串表示16进制的哈希值，coinbase交易没有，BLOCK数据',
  `b_vout_num_org` SMALLINT(5) UNSIGNED NOT NULL COMMENT '输入在来源交易中的顺序，coinbase交易没有，BLOCK数据',
  `b_script_sig`   VARCHAR(19000)       NOT NULL COMMENT '解锁脚本，字符串，存asm字段，coinbase交易没有，BLOCK数据',
  `b_txinwitness`  VARCHAR(2000)        NOT NULL COMMENT '隔离见证数据，字符串表示16进制，coinbase交易没有，BLOCK数据',
  `b_sequence`     INT(10) UNSIGNED     NOT NULL COMMENT '序列，默认0xffffffff（功能未启用，都是0xffffffff），BLOCK数据',
  `c_address_org`  VARCHAR(256)         NOT NULL COMMENT '输入的地址，通过blk_txid_org和blk_vout_num_org查询t_btc_vout表获得，coinbase交易没有，计算数据',
  `c_value_org`    BIGINT(20) UNSIGNED  NOT NULL COMMENT '输入的价值，单位聪，通过blk_txid_org和blk_vout_num_org查询t_btc_vout表获得，coinbase交易没有，计算数据',
  `b_time`         INT(10) UNSIGNED     NOT NULL COMMENT '时间戳，BLOCK数据',
  PRIMARY KEY (`id`),
  KEY `idx_b_txid` (`b_txid`),
  KEY `idx_b_block_height` (`b_block_height`),
  KEY `idx_b_time` (`b_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type VIn struct {
	ID               int64  `xorm:"id bigint autoincr pk"`
	BlockHeight      int64  `xorm:"block_height int notnull"`
	TxID             string `xorm:"tx_id char(64) notnull"`
	TxIDOrigin       string `xorm:"tx_id_origin char(64) notnull"`
	Timestamp        int64  `xorm:"timestamp int notnull"`
	Coinbase         string `xorm:"coinbase varchar(200) notnull"`
	Sequence         int64  `xorm:"sequence bigint notnull"`
	Number           int64  `xorm:"number int notnull"`
	VOutNumberOrigin int64  `xorm:"vout_num_origin smallint notnull"`
	AddressOrigin    string `xorm:"address_origin varchar(256) notnull"`
	ValueOrigin      int64  `xorm:"value_origin bigint notnull"`
	ScriptSignature  string `xorm:"script_sig text notnull"`
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
