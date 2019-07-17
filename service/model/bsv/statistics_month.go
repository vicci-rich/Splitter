package bsv

/*
DROP TABLE IF EXISTS `btc_month`;
CREATE TABLE `btc_month` (
  `id`                   INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `timestamp`            INT(10)          NOT NULL COMMENT '日期',
  `active_address_count` INT(10) UNSIGNED NOT NULL COMMENT '月活',
  `utxo_count`           INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'UTXO活跃率，数量',
  `utxo_address_count`   INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'UTXO活跃率，地址数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_time` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type StatisticsMonth struct {
	ID                 int64 `xorm:"id bigint autoincr pk"`
	Timestamp          int64 `xorm:"timestamp int notnull unique"`
	ActiveAddressCount int64 `xorm:"active_address_count int notnull default '0'"`
}

func (t StatisticsMonth) TableName() string {
	return tableName("statistics_month")
}
