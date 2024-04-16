package database

import (
	"database/sql"
	"fmt"

	"github.com/cloyop/tiniest/internal/types"
)

func InsertVisit(v *types.Visit) error {
	return db.QueryRow(`INSERT INTO visits (key,title,visit_from,url_visited,user_owner,qr,date_timestamp,date,hour ) VALUES (?,?,?,?,?,?,?,?,?)`, &v.Key, &v.Title, &v.Visit_from, &v.Url_visited, &v.User_owner, &v.Qr, &v.DateTimestamp, &v.Date, &v.Date_hour).Err()
}
func DeleteAllViewsFromKey(key string) error {
	row := db.QueryRow(`DELETE FROM visits WHERE key = ?`, key)
	return row.Err()
}
func DeleteAllViewFromUser(user_id string) error {
	row := db.QueryRow(`DELETE FROM visits WHERE user_owner = ?`, user_id)
	return row.Err()
}
func GetAllVisitsFromOwnerID(user_id string) (*[]types.Visit, error) {
	var visits []types.Visit
	query := `SELECT * FROM visits WHERE user_owner = ? ORDER BY date_timestamp DESC`
	rows, err := db.Query(query, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &visits, nil
		}
		return &visits, err
	}
	defer rows.Close()
	v := types.Visit{}
	for rows.Next() {
		rows.Scan(&v.Key, &v.Title, &v.Visit_from, &v.Url_visited, &v.User_owner, &v.Qr, &v.DateTimestamp, &v.Date, &v.Date_hour)
		visits = append(visits, v)
	}
	return &visits, nil
}
func GetVisitsFiltered(user_id, ext string) (*[]types.Visit, error) {
	var visits []types.Visit
	str := fmt.Sprintf("SELECT * FROM visits WHERE user_owner = '%v' %v ORDER BY date_timestamp DESC", user_id, ext)
	rows, err := db.Query(str)
	if err != nil {
		if err == sql.ErrNoRows {
			return &visits, nil
		}
		return &visits, err
	}
	v := types.Visit{}
	for rows.Next() {
		rows.Scan(&v.Key, &v.Title, &v.Visit_from, &v.Url_visited, &v.User_owner, &v.Qr, &v.DateTimestamp, &v.Date, &v.Date_hour)
		visits = append(visits, v)
	}
	return &visits, nil
}
func AddFilter(name, value string) string {
	return fmt.Sprintf(` AND %v="%v"`, name, value)
}
