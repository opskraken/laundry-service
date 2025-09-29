package userRepo

import "fmt"

func (ur *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
		username,
		email, 
		password, 
		status, 
		is_shop_owner
	) VALUES(
		$1, $2, $3, $4, $5
	) 
	RETURNING id
	`
	row := ur.db.QueryRow(query, user.UserName, user.Email, user.Password, user.Status, user.IsShopOwner)
	if row.Err() != nil {
		fmt.Println("error row", row.Err().Error())
		return nil, row.Err()
	}
	row.Scan(&user.Id)

	return &user, nil

}
