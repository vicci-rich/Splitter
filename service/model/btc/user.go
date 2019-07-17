package btc

type User struct {
	ID       int64  `xorm:"id bigint autoincr pk"`
	UserId   int64  `xorm:"user_id bigint notnull  default '0' unique"`
	Category string `xorm:"category varchar(20) notnull"`
}

func (t User) TableName() string {
	return tableName("user")
}
