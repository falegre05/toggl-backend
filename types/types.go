package types

import "github.com/golang-jwt/jwt"

// Question represents a quiz question with a body and a list of options
type Question struct {
	ID      int      `json:"id" db:"id"`
	Body    string   `json:"body" db:"body"`
	Options []Option `json:"options,omitempty"`
	UserID  int      `json:"user_id" db:"user_id"`
}

// Option represents an option in the quiz question
type Option struct {
	ID         int    `json:"id" db:"id"`
	Body       string `json:"body" db:"body"`
	Correct    bool   `json:"correct" db:"correct"`
	QuestionID int    `json:"questionId" db:"question_id"`
}

// User represents a user in the system
type User struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"-" db:"password_hash"`
}

// JWTClaims represents the claims in a JWT token
type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
