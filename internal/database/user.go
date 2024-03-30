package database

import (
	"database/sql"

	"github.com/cloyop/tiniest/internal/types"
)

func InsertUser(u *types.User) error {
	err := db.QueryRow(`INSERT INTO users (id,name,email,password,avatar_url,protected,verified,joined_at,links_count,max_links,email_last_update,name_last_update,password_last_update) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`, u.Id, u.Name, u.Email, u.Password, u.Avatar_url, u.Protected, u.Verified, u.Joined_Date, u.LinksCount, u.MaxLinksCount, 0, 0, 0).Err()
	return err
}

func GetUserByMail(mail string) (*types.User, bool, error) {
	var u types.User
	err := db.QueryRow(`SELECT * FROM users WHERE email = ?`, mail).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Avatar_url, &u.Protected, &u.Verified, &u.Joined_Date, &u.MaxLinksCount, &u.LinksCount, &u.Updates.Email_Last_Update, &u.Updates.Name_Last_Update, &u.Updates.Password_Last_Update)
	if err != nil {
		if err == sql.ErrNoRows {
			return &u, false, nil
		}
		return &u, false, err
	}
	u.Updates.FullFill()
	return &u, true, nil
}
func DeleteUser(mail string) error {
	return db.QueryRow(`DELETE FROM users WHERE email = ?`, mail).Err()
}
func UpdatePassword(password, id string) error {
	return db.QueryRow(`UPDATE users SET password=?, protected=? WHERE id=? RETURNING *`, password, true, id).Err()
}
func UpdateUser(u *types.User) error {
	return db.QueryRow(`UPDATE users SET email=?, name=?, password=?, protected=?,email_last_update=?,name_last_update=?,password_last_update=? WHERE id=?`, u.Email, u.Name, u.Password, u.Protected, u.Updates.Email_Last_Update, u.Updates.Name_Last_Update, u.Updates.Password_Last_Update, u.Id).Err()
}
func UpdateMaxLinks(u *types.User) error {
	return db.QueryRow(`UPDATE users SET max_links=? WHERE id=? `, u.MaxLinksCount, u.Id).Err()
}
func UpdateUserLinksCount(u *types.User) error {
	return db.QueryRow(`UPDATE users SET links_count=?  WHERE id=? `, u.LinksCount, u.Id).Err()
}
func GetAllUsers() (*[]types.User, error) {
	var users []types.User
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		if err == sql.ErrNoRows {
			return &users, nil
		}
		return &users, err
	}
	defer rows.Close()
	u := types.User{}
	for rows.Next() {
		rows.Scan(
			&u.Id,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Avatar_url,
			&u.Protected,
			&u.Verified,
			&u.Joined_Date,
			&u.LinksCount,
			&u.MaxLinksCount,
			&u.Updates.Email_Last_Update,
			&u.Updates.Name_Last_Update,
			&u.Updates.Password_Last_Update)
		users = append(users, u)
	}
	return &users, nil
}
