package postgres

import (
	"context"
	"time"

	"github.com/blawhi2435/shanjuku-backend/database/postgres"
	"github.com/blawhi2435/shanjuku-backend/domain"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/mapper/repository"
	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type activityName struct {
	ActivityName string    `structs:"activity_name"`
	Days         int       `structs:"days"`
	StartDate    time.Time `structs:"start_date"`
}

type activityPostgresRepository struct {
	db *gorm.DB
}

func ProvideActivityPostgresRepository(db *gorm.DB) domain.ActivityPostgreRepository {
	return &activityPostgresRepository{db}
}

func (a *activityPostgresRepository) InsertActivity(ctx context.Context, activity *domain.Activity) error {

	schemaActivity := repository.MappingActivityDomainToSchema(activity)
	orm := a.db.Model(&schemaActivity)
	res := orm.Select("id", "group_id", "creator_id", "activity_name", "days", "startDate").
		Create(&schemaActivity)

	return res.Error
}

func (a *activityPostgresRepository) QueryByID(ctx context.Context, activityID int64) (domain.Activity, error) {
	var schemaActivity postgres.Activity
	orm := a.db.Model(&schemaActivity)
	res := orm.Select("id", "group_id", "creator_id", "activity_name", "days", "startDate").
		Where("id = ?", activityID).
		Take(&schemaActivity)

	if res.Error == gorm.ErrRecordNotFound {
		return domain.Activity{}, cerror.ErrActivityNotExist
	}

	return repository.MappingActivitySchemaToDomain(&schemaActivity), res.Error
}

func (a *activityPostgresRepository) UpdateActivityName(ctx context.Context,
	activity *domain.Activity) (int64, error) {

	updateColumn := activityName{
		ActivityName: activity.ActivityName,
		Days:         activity.Days,
		StartDate:    activity.StartDate,
	}

	res := a.db.Model(&postgres.Activity{}).
		Where("id = ?", activity.ID).
		Updates(structs.Map(updateColumn))

	return res.RowsAffected, res.Error
}

func (a *activityPostgresRepository) DeleteActivityWithTx(ctx context.Context, tx any,
	activityID int64) error {

	var orm *gorm.DB
	if tx == nil {
		orm = a.db
	} else {
		orm = tx.(*gorm.DB)
	}

	res := orm.Delete(&postgres.Activity{}, activityID)

	return res.Error
}
