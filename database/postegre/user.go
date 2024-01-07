package postegre

import "time"

type User struct {
	ID          int64     `gorm:"type:bigint NOT NULL;primary_key"`
	Account     string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	Password    string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	Name        string    `gorm:"type:varchar(16) NOT NULL DEFAULT ''"`
	Avatar      string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	CreatedDate time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	Groups    []Group    `gorm:"many2many:user_group;"`
	Schedules []Schedule `gorm:"many2many:user_schedule;"`
}
