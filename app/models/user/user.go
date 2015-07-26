package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ShaneBurkhart/PlanSource/config/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"sort"
	"time"
)

type User struct {
	Id             int
	FirstName      string
	LastName       string
	Company        string
	Email          string
	Password       string
	PasswordDigest string
	Role           string
	LastSeen       time.Time
	UpdatedAt      time.Time
	CreatedAt      time.Time
	Errors         []error
}

func All() []*User {
	users := make([]*User, 0)
	rows, err := db.DB.Query(findAllSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := &User{}
		err := scan(u, rows)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return users
		}
		//TODO Logging
		log.Fatal(err)
	}

	return users
}

// Determines the column to sort by and the order.
// Return vals: column, order
func sortingKeys(c string, o string) (string, string) {
	// Since we can't do sql params in an order by clause, we have to white list
	// the columns we want and the order.
	// Some defaults if the value isn't in white list.
	var column string = "role_id"
	var order string = "ASC"
	if c == "role_id" || c == "name" || c == "email" || c == "company" ||
		c == "created_at" || c == "last_seen" {
		column = c
	}
	if o == "ASC" || o == "DESC" {
		order = o
	}
	return column, order
}

func AllSorted(c string, o string) []*User {
	users := make([]*User, 0)
	column, order := sortingKeys(c, o)
	query := fmt.Sprintf("%s ORDER BY %s %s", findAllSQL, column, order)
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := &User{}
		err := scan(u, rows)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return users
		}
		//TODO Logging
		log.Fatal(err)
	}

	if column == "email" {
		i := sort.Interface(byEmailDomain(users))
		if order == "DESC" {
			i = sort.Reverse(byEmailDomain(users))
		}
		sort.Sort(i)
	}
	if column == "company" {
		i := sort.Interface(byCompany(users))
		if order == "DESC" {
			i = sort.Reverse(byCompany(users))
		}
		sort.Sort(i)
	}

	return users
}

func Find(id int) *User {
	u := &User{}
	err := scan(u, db.DB.QueryRow(findByIdSQL, id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		//TODO Logging
		log.Fatal(err)
	}
	return u

}

func (u *User) Login() bool {
	err := scan(u, db.DB.QueryRow(findByEmailSQL, u.Email))
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		//TODO Logging
		// Some kind of sql error
		log.Fatal(err)
	}
	return u.comparePassword(u.Password)
}

func (u *User) UpdateLastSeen() {
	_, err := db.DB.Exec(updateLastSeenSQL, u.Id)
	if err != nil {
		log.Print(err)
	}
}

func (u *User) Delete() bool {
	// We are ignoring the result which only tells us the number of rows affected.
	_, err := db.DB.Exec(
		deleteUserSQL,
		u.Id,
	)

	if err != nil {
		log.Print(err)
		return false
	}
	log.Print("Deleted ", u)

	u.Id = 0
	return true
}

func (u *User) Save() bool {
	u.validate()
	if u.HasErrors() {
		return false
	}

	if u.Id <= 0 {
		return u.create()
	} else {
		return u.update()
	}
}

func (u *User) create() bool {
	var id int
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO don't crash server
		log.Fatal(err)
	}

	err = db.DB.QueryRow(
		insertUserSQL,
		u.FirstName,
		u.LastName,
		u.Company,
		u.Email,
		passwordDigest,
		u.Role,
	).Scan(&id)
	if err != nil {
		// Most likely this error is due to index constraints not being met.
		// In this case, it is most likely that the email is not unique.
		log.Print(err)
		// TODO Error message
		return false
	}
	// Erase password since we no longer need it.
	u.Password = ""
	u.Id = id
	u.PasswordDigest = string(passwordDigest)

	log.Print("Added ", u)
	return true
}

func (u *User) update() bool {
	if u.Password != "" {
		passwordDigest, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		u.PasswordDigest = string(passwordDigest)
		if err != nil {
			// TODO don't crash server
			log.Fatal(err)
		}
	}

	_, err := db.DB.Exec(
		updateUserSQL,
		u.Id,
		u.FirstName,
		u.LastName,
		u.Company,
		u.Email,
		u.PasswordDigest,
		u.Role,
	)
	if err != nil {
		// Most likely this error is due to index constraints not being met.
		// In this case, it is most likely that the email is not unique.
		log.Print(err)
		// TODO Error message
		return false
	}
	// Erase password since we no longer need it.
	u.Password = ""

	log.Print("Updated", u)
	return true
}

func (u *User) UpdatePassword(old_pass string, new_pass string, new_pass_conf string) {
	if new_pass == "" {
		return
	}
	if new_pass != new_pass_conf {
		u.addError("Your new password and it's confirmation don't match.")
		return
	}
	if !u.comparePassword(old_pass) {
		u.addError("Your old password isn't correct.")
		return
	}
	u.Password = new_pass
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) IsManager() bool {
	return u.Role == "manager"
}

func (u *User) IsViewer() bool {
	return u.Role == "viewer"
}

func (u *User) HasErrors() bool {
	return u.Errors != nil && len(u.Errors) > 0
}

func (u *User) addError(message string) {
	if len(message) > 0 {
		u.Errors = append(u.Errors, errors.New(message))
	}
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scan(u *User, row scanner) error {
	return row.Scan(
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Company,
		&u.Email,
		&u.PasswordDigest,
		&u.Role,
		&u.LastSeen,
		&u.UpdatedAt,
		&u.CreatedAt,
	)
}

func (u *User) validate() {
	// Check email
	// TODO Check for duplicate email
	if matched, _ := regexp.MatchString("^\\S+@\\S+$", u.Email); !matched {
		u.addError("Email is invalid.")
	}

	// Check password
	if u.Id == 0 || len(u.Password) > 0 {
		// We have to have a password for first save or if the password is set and
		// id is not zero, we are updating the password so it needs to be validated.
		u.addError(validatePassword(u.Password))
	}
}

func validatePassword(password string) string {
	// TODO More password validation
	if len(password) == 0 {
		return "Password cannot be blank."
	}
	return ""
}

func (u *User) comparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
