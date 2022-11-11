package repository

import (
	"context"
	"database/sql"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx sql.Tx)
	Update(ctx context.Context, tx sql.Tx)
	Delete(ctx context.Context, tx sql.Tx)
	FindOne(ctx context.Context, tx sql.Tx)
	FindAll(ctx context.Context, tx sql.Tx)
}
