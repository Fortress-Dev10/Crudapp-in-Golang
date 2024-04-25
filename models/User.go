package models

type user struct {
	ID       uint	`json:"id"`
	Username string	`json:"username"`
	Email    string	`json:"email"`
	Password string	`json:"password"`
}

func (u *user) TableName() string {
	return "users"
}

func (u *user) Create() error {
	//insert into users (username, email, password) values (u.Username, u.Email, u.Password)
	return nil
}


func (u *user) Update() error {
	//update users set username = u.Username, email = u.Email, password = u.Password where id = u.ID
	return nil
}

func (u *user) Delete() error {
	//delete from users where id = u.ID
	return nil
}


