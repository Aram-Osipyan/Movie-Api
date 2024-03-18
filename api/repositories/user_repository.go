package repositories

import "github.com/Movie-Api/models"

type UserRepository struct {
	BaseRepository
}

type UserNotFoundError struct{}

func (u *UserNotFoundError) Error() string {
	return "user not found error"
}

func (user_repository *UserRepository) Create(username, password, role string) (*models.User, error) {
	db, err := user_repository.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("INSERT INTO users (username, password, role) VALUES($1, $2, $3) RETURNING id;", username, password, role)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	new_user := models.User{Username: username, Password: password, Role: role}

	for rows.Next() {
		err = rows.Scan(&new_user.Id)
		break
	}

	return &new_user, err
}

func (user_repository *UserRepository) Find(username, password string) (*models.User, error) {
	db, err := user_repository.GetConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, username, password, role FROM users WHERE username = $1 AND password = $2", username, password)

	defer db.Close()

	if err != nil {
		return nil, &UserNotFoundError{}
	}
	user := models.User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
		break
	}

	return &user, err
}
