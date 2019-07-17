package bch

/*
DROP TABLE IF EXISTS `btc_vout`;
CREATE TABLE `btc_vout` (
  `id`              INT(10) UNSIGNED     NOT NULL AUTO_INCREMENT,
  `b_txid`          CHAR(64)             NOT NULL COMMENT '交易id，固定256bit，返回值用64位字符串表示16进制的哈希值，BLOCK数据',
  `r_number`        SMALLINT(5) UNSIGNED NOT NULL COMMENT '输出在交易中的顺序得，RPC数据',
  `b_block_height`  INT(10) UNSIGNED     NOT NULL COMMENT '区块高度，BLOCK数据',
  `b_value`         BIGINT(20) UNSIGNED  NOT NULL COMMENT '输出价值，单位聪，BLOCK数据',
  `r_type`          VARCHAR(256)         NOT NULL COMMENT '输出地址类型，字符串，RPC数据',
  `b_address`       VARCHAR(256)         NOT NULL COMMENT  '输出地址，字符串，BLOCK数据',
  `b_script_pubkey` VARCHAR(10000)       NOT NULL COMMENT '锁定脚本，字符串，存hex字段（损失一定可读性，可以自己转换OP_code），BLOCK数据',
  `r_reqsigs`       SMALLINT(5) UNSIGNED NOT NULL COMMENT '解锁脚本需要的签名数，一般为1，多重签名脚本可能大于1，RPC数据',
  `c_isused`        TINYINT(4) UNSIGNED  NOT NULL COMMENT '是否被使用 通过查询vin确定',
  `b_time`          INT(10) UNSIGNED     NOT NULL COMMENT '时间戳，BLOCK数据',
  `c_iscoinbase`    TINYINT(4) UNSIGNED  NOT NULL COMMENT '是否coinbase',
  PRIMARY KEY (`id`),
  KEY `idx_b_txid_r_number` (`b_txid`, `r_number`),
  KEY `idx_b_block_height` (`b_block_height`),
  KEY `idx_b_time` (`b_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type VOut struct {
	ID                 int64  `xorm:"id bigint autoincr pk"`
	TxID               string `xorm:"tx_id char(64) notnull"`
	BlockHeight        int64  `xorm:"block_height int notnull"`
	Value              uint64 `xorm:"value bigint notnull"`
	Address            string `xorm:"address varchar(256) notnull"`
	Timestamp          int64  `xorm:"timestamp int notnull"`
	ScriptPublicKey    string `xorm:"script_pubkey text notnull"`
	Type               string `xorm:"type varchar(256) notnull"`
	RequiredSignatures int64  `xorm:"required_signatures smallint notnull"`
	Number             int64  `xorm:"number smallint notnull"`
	IsUsed             int64  `xorm:"is_used tinyint notnull"`
	IsCoinbase         int64  `xorm:"is_coinbase tinyint notnull"`
}

func (t VOut) TableName() string {
	return tableName("vout")
}
