package postegre

import "time"

type Schedule struct {
	ID          int64     `gorm:"type:bigint NOT NULL;primary_key"`
	ActivityID  int64     `gorm:"type:bigint NOT NULL"`
	Name        string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	Comment     string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	StartDate   time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	EndDate     time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedDate time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	Activity Activity `gorm:"foreignkey:ActivityID"`
	Users    []User   `gorm:"many2many:user_schedule;"`
}
