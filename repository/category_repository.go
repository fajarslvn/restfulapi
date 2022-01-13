package repository

import (
	"context"
	"database/sql"

	"github.com/fajarslvn/restfulapi/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

/* Interface in Go should be like this!
type Saver interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
}

type Updater interface {
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
}

type Deleter interface {
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
}

type Finder interface {
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
*/
