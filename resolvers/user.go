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

// UserLogin performs user authentication by checking the provided credentials.
// If the user exists and the password is correct, it generates a JWT token for the user.
func UserLogin(args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		Name     string
		Password string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Check if the user actually exists
	var user types.User
	if err := database.GetDBConnection().Get(&user, "SELECT * FROM users WHERE name = ?", input.Name); err != nil {
		return nil, err
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, errors.New("user or password incorrect")
	}

	// Generate a JWT token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.JWTClaims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{IssuedAt: jwt.TimeFunc().Unix()},
	})

	// Sign the token with the secret key and return the signed token
	return token.SignedString([]byte(security.GetSecretKey()))
}

// AddUser creates a new user with the provided username and password.
func AddUser(args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		Name     string
		Password string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Insert the new user into the database
	result, err := database.GetDBConnection().Exec("INSERT INTO users (name, password_hash) VALUES (?, ?)", input.Name, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted user
	id, _ := result.LastInsertId()

	// Return the newly added user
	return getUserByID(int(id))
}

// getUserByID retrieves a user by their ID from the database.
func getUserByID(id int) (types.User, error) {
	var user types.User
	err := database.GetDBConnection().Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return types.User{}, err
	}
	return user, err
}
