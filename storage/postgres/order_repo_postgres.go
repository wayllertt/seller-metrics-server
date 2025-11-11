package postgres

import (
	"context"
	"database/sql"

	"seller-metrics-server/internal/domain"
)

type OrderRepoPostgres struct {
	db *sql.DB
}

func NewOrderRepoPostgres(db *sql.DB) *OrderRepoPostgres {
	return &OrderRepoPostgres{db: db}
}

func (r *OrderRepoPostgres) GetBySellerID(ctx context.Context, sellerID int) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, seller_id, created_at, promised_at, delivered_at
        FROM orders
        WHERE seller_id = $1
    `, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.Order
	for rows.Next() {
		var o domain.Order
		var deliveredAt sql.NullTime

		if err := rows.Scan(
			&o.ID,
			&o.SellerID,
			&o.CreatedAt,
			&o.PromisedAt,
			&deliveredAt,
		); err != nil {
			return nil, err
		}

		if deliveredAt.Valid {
			o.DeliveredAt = &deliveredAt.Time
		}

		res = append(res, o)
	}
	return res, rows.Err()
}
