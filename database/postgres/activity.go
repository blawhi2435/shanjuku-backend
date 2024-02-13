package postgres

import "time"

type Activity struct {
	ID           int64     `gorm:"type:bigint NOT NULL;primary_key"`
	GroupID      int64     `gorm:"type:bigint NOT NULL"`
	CreatorID    int64     `gorm:"type:bigint NOT NULL"`
	ActivityName string    `gorm:"type:varchar(64) NOT NULL DEFAULT ''"`
	Days         int       `gorm:"type:int NOT NULL DEFAULT 1"`
	StartDate    time.Time `gorm:"type:date NOT NULL DEFAULT CURRENT_DATE"`
	CreatedDate  time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	Group Group `gorm:"foreignkey:GroupID"`
	User  User  `gorm:"foreignkey:CreatorID"`
}
