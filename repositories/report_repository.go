package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailyReport(date time.Time) (*models.DailyReport, error) {
	start := date.Format("2006-01-02") + " 00:00:00"
	end := date.Format("2006-01-02") + " 23:59:59"

	var report models.DailyReport

	err := repo.db.QueryRow(`
		SELECT
			COALESCE(SUM(total_amount), 0),
			COUNT(id)
		FROM transactions
		WHERE created_at BETWEEN $1 AND $2
	`, start, end).Scan(
		&report.TotalRevenue,
		&report.TotalTransaksi,
	)

	if err != nil {
		return nil, err
	}

	var best models.BestSellingProduct

	err = repo.db.QueryRow(`
		SELECT
			p.name,
			SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`, start, end).Scan(
		&best.Nama,
		&best.QtyTerjual,
	)

	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = best

	return &report, nil
}
