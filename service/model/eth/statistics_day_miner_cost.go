package eth

/*
DROP TABLE IF EXISTS `eth_cost`;
CREATE TABLE `eth_cost` (
  `id`                   INT(10) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `time`                 VARCHAR(20)         NOT NULL COMMENT '日期',
  `energy_c_per_day`     DOUBLE              NOT NULL DEFAULT '0' COMMENT '能源消耗（元/天/G/）固定值',
  `cost`                 DOUBLE              NOT NULL DEFAULT '0' COMMENT '挖矿成本（元/天/G）固定值',
  `eth_output`           DOUBLE              NOT NULL DEFAULT '0' COMMENT '以太币产量（个/天/G）1/全网算力*24*4*60',
  `energy_c_per_eth`     DOUBLE              NOT NULL DEFAULT '0' COMMENT '单虚拟货币能源消耗（元/天/G/个）energy_c_per_day/eth_output',
  `income`               DOUBLE              NOT NULL DEFAULT '0' COMMENT '挖矿收益（元/天/G）eth_output*price',
  `profit`               DOUBLE              NOT NULL DEFAULT '0' COMMENT '挖矿纯收益（元/天/G）income-cost',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_time` (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `currency_exchange_rates` (
  `id`                   BIGINT        NOT NULL AUTO_INCREMENT,
  `timestamp`            INT           NOT NULL,
  `usd`                  DOUBLE        NOT NULL,
  `cny`                  DOUBLE        NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_timestamp` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type StatisticsDayMinerCost struct {
	ID                     int64   `xorm:"id bigint autoincr pk"`
	Timestamp              int64   `xorm:"timestamp int notnull unique"`
	Cost                   float64 `xorm:"cost decimal(38,4) notnull default '0'"`
	TotalEnergyConsumption float64 `xorm:"total_energy_consumption decimal(38,4) notnull default '0'"`
	UnitEnergyConsumption  float64 `xorm:"unit_energy_consumption decimal(38,4) notnull default '0'"`
	Output                 float64 `xorm:"output decimal(38,4) notnull default '0'"`
	Income                 float64 `xorm:"income decimal(38,4) notnull default '0'"`
	Profit                 float64 `xorm:"profit decimal(38,4) notnull default '0'"`
}

func (t StatisticsDayMinerCost) TableName() string {
	return tableName("statistics_day_miner_cost")
}
