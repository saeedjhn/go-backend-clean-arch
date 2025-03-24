package respository

import (
	"github.com/saeedjhn/go-backend-clean-arch/pkg/persistance/db/mysql"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) *DB {
	return &DB{conn: conn}
}

// func (d DB) InsertEvent(ctx context.Context, evt outbox.Event) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (d DB) UpdatePublished(ctx context.Context, eventIDs []types.ID, publishedAt time.Time) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (d DB) UpdateUnpublished(ctx context.Context, eventIDs []types.ID, lastRetriedAt time.Time) error {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (d DB) UnpublishedCount(ctx context.Context, retryThreshold int64) (int64, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
//
// func (d DB) GetUnPublished(ctx context.Context, offset, limit, retryThreshold int) ([]outbox.Event, error) {
// 	// TODO implement me
// 	panic("implement me")
// }
