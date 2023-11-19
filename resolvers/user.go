package resolvers

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"github.com/togglhire/backend-homework/database"
	"github.com/togglhire/backend-homework/security"
	"github.com/togglhire/backend-homework/types"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(args map[string]interface{}) (interface{}, error) {
	var input struct {
		Name     string
		Password string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Check if user actually exists
	var user types.User
	if err := database.GetDBConnection().Get(&user, "SELECT * FROM users WHERE name = ?", input.Name); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, errors.New("user or password incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.JWTClaims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{IssuedAt: jwt.TimeFunc().Unix()},
	})

	return token.SignedString([]byte(security.GetSecretKey()))
}

func AddUser(args map[string]interface{}) (interface{}, error) {
	var input struct {
		Name     string
		Password string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Hashing password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Insert the new question into the database
	result, err := database.GetDBConnection().Exec("INSERT INTO users (name, password_hash) VALUES (?, ?)", input.Name, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted question
	id, _ := result.LastInsertId()

	// Return the newly added question
	return getUserByID(int(id))
}

func getUserByID(id int) (types.User, error) {
	var user types.User
	err := database.GetDBConnection().Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return types.User{}, err
	}
	return user, err
}
