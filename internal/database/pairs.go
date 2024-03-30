package database

import (
	"database/sql"
	"fmt"

	"github.com/cloyop/tiniest/internal/types"
)

func InsertPair(p *types.PairLinks) error {
	row := db.QueryRow(`INSERT INTO links (key,short,long,user_id,date,protected,password,deleted,title) VALUES (?,?,?,?,?,?,?,?,?)`, p.Key, p.Short, p.Long, p.User_id, p.Date, p.Protected, p.Password, p.Deleted, p.Title)
	return row.Err()
}
func MigratePairOwnership(pairKey string, u *types.User) {
	var p types.PairLinks
	p.Key = pairKey
	err := db.QueryRow("UPDATE links SET user_id=? WHERE key=? RETURNING short", u.Id, pairKey).Scan(&p.Short)
	if err == nil {
		p.GenerateQR()
	}
}
func LockPair(p, k, id string) (*types.PairLinks, error) {
	var pair types.PairLinks
	err := db.QueryRow("UPDATE links SET protected = ?,password = ? WHERE key = ? AND user_id = ? RETURNING * ", true, p, k, id).Scan(&pair.Key, &pair.Short, &pair.Long, &pair.User_id, &pair.Date, &pair.Protected, &pair.Password, &pair.Deleted, &pair.Title)
	return &pair, err
}
func DeletePair(sh, id string) (*types.PairLinks, error) {
	var pair types.PairLinks
	err := db.QueryRow(`DELETE FROM links WHERE key = ? AND user_id = ? RETURNING *`, sh, id).Scan(&pair.Key, &pair.Short, &pair.Long, &pair.User_id, &pair.Date, &pair.Protected, &pair.Password, &pair.Deleted, &pair.Title)
	pair.DeleteQR()
	return &pair, err
}
func DeleteAllPairFromUser(user_id string) error {
	row := db.QueryRow(`DELETE FROM links WHERE user_id = ?`, user_id)
	return row.Err()
}
func DeletePairPassword(key, id string) (types.PairLinks, error) {
	var pair types.PairLinks
	err := db.QueryRow("UPDATE links SET protected = ?,password = ? WHERE key = ? AND user_id = ? RETURNING * ", false, "", key, id).Scan(&pair.Key, &pair.Short, &pair.Long, &pair.User_id, &pair.Date, &pair.Protected, &pair.Password, &pair.Deleted, &pair.Title)
	return pair, err
}
func GetPairPassword(key, id string) (string, error) {
	var pwd string
	err := db.QueryRow(`SELECT password FROM links WHERE key = ? AND user_id=?`, key, id).Scan(&pwd)
	return pwd, err
}
func GetPair(key string) (*types.PairLinks, bool, error) {
	var p types.PairLinks
	err := db.QueryRow(`SELECT * FROM links WHERE key = ?`, key).Scan(&p.Key, &p.Short, &p.Long, &p.User_id, &p.Date, &p.Protected, &p.Password, &p.Deleted, &p.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return &p, false, nil
		}
		return &p, false, err
	}
	return &p, true, err
}
func GetPairWithId(key, id string) (*types.PairLinks, bool, error) {
	var p types.PairLinks
	err := db.QueryRow(`SELECT * FROM links WHERE key = ? AND user_id= ?`, key, id).Scan(&p.Key, &p.Short, &p.Long, &p.User_id, &p.Date, &p.Protected, &p.Password, &p.Deleted, &p.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return &p, false, nil
		}
		return &p, false, err
	}
	return &p, true, err
}
func GetAllPairsByID(user_id string) (*[]types.PairLinks, error) {
	var pairs []types.PairLinks
	query := `SELECT * FROM links WHERE user_id = ?`
	rows, err := db.Query(query, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pairs, nil
		}
		return &pairs, err
	}
	defer rows.Close()
	p := types.PairLinks{}
	for rows.Next() {
		rows.Scan(&p.Key, &p.Short, &p.Long, &p.User_id, &p.Date, &p.Protected, &p.Password, &p.Deleted, &p.Title)
		pairs = append(pairs, p)
	}
	return &pairs, nil
}
func GetAllPairsByKeys(strKeys string) (*[]types.PairLinks, error) {
	var pairs []types.PairLinks
	query := fmt.Sprintf(`SELECT * FROM links WHERE key in (%v)`, strKeys)
	rows, err := db.Query(query)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pairs, nil
		}
		return &pairs, err
	}
	defer rows.Close()
	p := types.PairLinks{}
	for rows.Next() {
		rows.Scan(&p.Key, &p.Short, &p.Long, &p.User_id, &p.Date, &p.Protected, &p.Password, &p.Deleted, &p.Title)
		pairs = append(pairs, p)
	}
	return &pairs, nil
}
