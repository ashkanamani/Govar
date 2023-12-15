package data

import "time"

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}


// Create a new user, save user info into the database
func (user *User) Create() error {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.ID, &user.UUID, &user.CreatedAt, )
	return err
}

// Delete user from database
func (user *User) Delete() error {
	statement := "delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return err
}

// Update user information in the database
func (user *User) Update() error {
	statement := "update users set name = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email)
	return err
}

// Get all users in the database and returns it
func Users() ([]User, error) {
	var users []User
	rows, err := Db.Query("select * from users")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}

// Get a single user given the email
func UserByEmail(email string) (User, error) {
	user := User{}
	err := Db.QueryRow("select * from users where email=$1", email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

// Get a single user given the UUID
func UserByUUID(UUID string) (User, error) {
	user := User{}
	err := Db.QueryRow("select * from users where uuid=$1", UUID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err

}