package domain

import (
	"context"
	"time"
)

type Activity struct {
	ID           int64     `json:"id"`
	GroupID      int64     `json:"group_id"`
	CreatorID    int64     `json:"creator_id"`
	ActivityName string    `json:"activity_name"`
	Days         int       `json:"days"`
	StartDate    time.Time `json:"start_date"`
}

type ActivityPostgreRepository interface {
	InsertActivity(ctx context.Context, activity *Activity) error
	QueryByID(ctx context.Context, activityID int64) (Activity, error)
	UpdateActivityName(ctx context.Context, activity *Activity) (int64, error)
	DeleteActivityWithTx(ctx context.Context, tx any, activityID int64) error
}

type ActivityUsecase interface {
	CreateActivity(ctx context.Context, activity *Activity) (Activity, error)
	QueryByID(ctx context.Context, activityID int64) (Activity, error)
	UpdateActivityName(ctx context.Context, activity *Activity) (Activity, error)
	DeleteActivity(ctx context.Context, activityID int64) error
	IsBelongToUser(ctx context.Context, activityID, userID int64) (bool, error)
	IsActivityExist(ctx context.Context, activityID int64) (bool, error)
}
