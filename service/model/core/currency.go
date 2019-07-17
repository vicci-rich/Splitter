package core

import (
	"time"
)

type Currency struct {
	ID          int64     `xorm:"id bigint autoincr pk"`
	CID         string    `xorm:"cid varchar(255) notnull unique"`
	Name        string    `xorm:"name varchar(255) notnull"`
	Symbol      string    `xorm:"symbol varchar(255) notnull"`
	Description string    `xorm:"description text notnull"`
	ReleaseTime int       `xorm:"release_time int notnull"`
	Site        string    `xorm:"site varchar(255) notnull"`
	Type        int       `xorm:"type tinyint notnull"`
	ParentCID   string    `xorm:"parent_cid varchar(255) notnull"`
	Rank        int       `xorm:"rank int notnull"`
	CreatedTime time.Time `xorm:"created_time created notnull"`
	UpdatedTime time.Time `xorm:"updated_time updated notnull"`
	DeletedTime time.Time `xorm:"deleted_time deleted default null"`
}

func (t Currency) TableName() string {
	return "currency"
}
