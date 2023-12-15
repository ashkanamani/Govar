package data

import "time"

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt time.Time
}


// Create a new session for an existing user
func (user *User)CreateSession() (Session, error) {
	var session Session
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return session, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return session, err
}

// Get the session for an existing user
func (user *User) Session() (Session, error) {
	var session Session
	var err error
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.ID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return session, err
}

// Check if session is valid in the database
func (session *Session) Check() (bool, error) {
	var valid bool
	err := Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.UUID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		valid = false
		return valid, err
	}
	if session.ID != 0 {
		valid = true
	}
	return valid, err
}

// Delete session from database
func (session *Session) DeleteByUUID() error {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.UUID)
	return err
}

// Get the user from the session
func (session *Session) User() (User, error) {
	user := User{}
	err := Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

// Delete all sessions from database
func SessionDeleteAll() error {
	statement := "delete from sessions"
	_, err := Db.Exec(statement)
	return err
}
