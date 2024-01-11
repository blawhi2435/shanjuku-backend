package postgres

import "time"

type Group struct {
	ID          int64     `gorm:"type:bigint NOT NULL;primary_key"`
	GroupName   string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	CreatedDate time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	Users []User `gorm:"many2many:user_group;"`
}
