package databse

import "app/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query(
		"SELECT id, first_name, last_name FROM users WHERE id = ?", identifier,
	)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var fitstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &fitstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = fitstName
	user.LastName = lastName
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	var users domain.Users

	for rows.Next() {
		var id int
		var fitstName string
		var lastName string
		if err = row.Scan(&id, &fitstName, &lastName); err != nil {
			continue
		}

		user := domain.User{
			ID: id,
			FirstName: firstName,
			LastName: lastName,
		}
		users = append(users, user)
	}
	return
}