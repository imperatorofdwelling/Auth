package users

import "github.com/jmoiron/sqlx"

func ScanRow(row *sqlx.Row) (*Model, error) {
	model := new(Model)
	err := row.StructScan(model)

	if err != nil {
		return nil, err
	}
	return model, nil
}
