package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/togglhire/backend-homework/database"
	"github.com/togglhire/backend-homework/security"
	"github.com/togglhire/backend-homework/types"
)

// GetAllQuestions retrieves all questions for the current user along with their options.
func GetAllQuestions(ctx context.Context) (interface{}, error) {
	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Retrieve all questions for the current user
	var questions []types.Question
	err := database.GetDBConnection().Select(&questions, "SELECT * FROM questions WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	// Retrieve options for each question
	for i := range questions {
		if err := database.GetDBConnection().Select(&questions[i].Options, "SELECT * FROM options WHERE question_id = ?", questions[i].ID); err != nil {
			return nil, err
		}
	}

	return questions, nil
}

// GetQuestionByID retrieves a specific question by its ID for the current user.
func GetQuestionByID(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Extract the ID from the input arguments
	id, ok := args["id"].(int)
	if !ok {
		return nil, errors.New("could not parse id")
	}

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Call the internal function to retrieve the question by ID
	return getQuestionByID(id, *userID)
}

// getQuestionByID is an internal function that retrieves a question by its ID and user ID.
func getQuestionByID(id, userID int) (types.Question, error) {
	var question types.Question
	if err := database.GetDBConnection().Get(&question, "SELECT * FROM questions WHERE id = ? AND user_id = ?", id, userID); err != nil {
		return types.Question{}, err
	}

	// Retrieve options for the question
	if err := database.GetDBConnection().Select(&question.Options, "SELECT * FROM options WHERE question_id = ?", question.ID); err != nil {
		return types.Question{}, err
	}
	return question, nil
}

// AddQuestion adds a new question along with its options for the current user.
func AddQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		Body    string
		Options []struct {
			Body    string
			Correct bool
		}
	}
	mapstructure.Decode(args, &input)

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Begin a new database transaction
	var tx, err = database.GetDBConnection().Begin()
	if err != nil {
		return nil, err
	}

	// Insert the new question into the database
	result, err := tx.Exec("INSERT INTO questions (body, user_id) VALUES (?, ?)", input.Body, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get the ID of the newly inserted question
	questionID, _ := result.LastInsertId()

	// Insert options for the question if they exist
	if len(input.Options) > 0 {
		stmt, err := tx.Prepare("INSERT INTO options (body, correct, question_id) VALUES (?, ?, ?)")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		for _, option := range input.Options {
			_, err := stmt.Exec(option.Body, option.Correct, questionID)
			if err != nil {
				return nil, err
			}
		}
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Return the newly added question
	return getQuestionByID(int(questionID), *userID)
}

// UpdateQuestion updates the details of an existing question for the current user.
func UpdateQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		ID   int
		Body string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Update the question in the database
	_, err := database.GetDBConnection().Exec("UPDATE questions SET body = ? WHERE id = ? AND user_id = ?", input.Body, input.ID, userID)
	if err != nil {
		return nil, err
	}

	// Return the updated question
	return getQuestionByID(input.ID, *userID)
}

// DeleteQuestion deletes an existing question for the current user.
func DeleteQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	id, ok := args["id"].(int)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'id' parameter")
	}

	var userID = security.GetUserID(ctx)

	// Update the question in the database
	result, err := database.GetDBConnection().Exec("DELETE FROM questions WHERE id = ? AND user_id = ?", id, userID)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("question with id %d does not exist", id)
	}

	return true, nil
}
