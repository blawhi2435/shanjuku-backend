package postegres

import "time"

type Activity struct {
	ID          int64     `gorm:"type:bigint NOT NULL;primary_key"`
	GroupID     int64     `gorm:"type:bigint NOT NULL"`
	CreatorID   int64     `gorm:"type:bigint NOT NULL"`
	Name        string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	StartDate   time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	EndDate     time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedDate time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	Group Group `gorm:"foreignkey:GroupID"`
	User  User  `gorm:"foreignkey:CreatorID"`
}
