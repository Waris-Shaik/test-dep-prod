package user

import (
	"database/sql"
	"fmt"
	"log"
	"test-dep-prod/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ? OR username = ?", email, email)
	if err != nil {
		log.Println("error in QUERY:", err)
		return nil, err
	}

	user := new(types.User)

	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			log.Println("error in rows.next()", err)
			return nil, err
		}
	}

	if user.ID == 0 {
		log.Println("user not found", user)
		return nil, fmt.Errorf("user not found")
	}

	log.Println("user found:", user)
	return user, nil

}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		log.Println("error while scanning row_into_user :", err)
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Println("error in QUERY:", err)
		return nil, err
	}

	user := new(types.User)

	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			log.Println("error in rows.next()", err)
			return nil, err
		}
	}

	if user.ID == 0 {
		log.Println("user not found", user)
		return nil, fmt.Errorf("user not found")
	}

	log.Println("user found:", user)
	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	result, err := s.db.Exec("INSERT INTO users (first_name, last_name, username, email, password) VALUES (?,?,?,?,?)", user.FirstName, user.LastName, user.UserName, user.Email, user.Password)
	if err != nil {
		log.Println("error in QUERY:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("error in generating lastInsertID()", err)
		return nil
	}

	fmt.Println("userID is:", id)

	return nil
}

func (s *Store) GetAllUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("error in QUERY:", err)
		return nil, err
	}

	users := make([]*types.User, 0)

	for rows.Next() {
		user, err := scanRowIntoUser(rows)
		if err != nil {
			log.Println("Error in rows.next", err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
