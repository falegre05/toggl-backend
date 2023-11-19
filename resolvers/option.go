package resolvers

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/togglhire/backend-homework/database"
	"github.com/togglhire/backend-homework/security"
	"github.com/togglhire/backend-homework/types"
)

// AddOptionToQuestion adds a new option to an existing question.
// It validates the input, checks if the question exists, and inserts the new option into the database.
func AddOptionToQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		QuestionID int
		Body       string
		Correct    bool
	}

	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Check if the question exists and if the user has access to it
	_, err := getQuestionByID(input.QuestionID, *userID)
	if err != nil {
		return nil, fmt.Errorf("question with id %d does not exist", input.QuestionID)
	}

	// Insert the new option into the database
	result, err := database.GetDBConnection().Exec("INSERT INTO options (question_id, body, correct) VALUES (?, ?, ?)", input.QuestionID, input.Body, input.Correct)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted option
	id, _ := result.LastInsertId()

	return getOptionByID(int(id))
}

// UpdateOption updates the details of an existing option.
// It validates the input, checks if the option exists, and updates the option in the database.
func UpdateOption(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Decode input arguments into a structured input object
	var input struct {
		ID      int
		Body    string
		Correct bool
	}

	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Check if the option exists and if the user has access to it
	_, err := getQuestionByOptionID(input.ID, *userID)
	if err != nil {
		return nil, fmt.Errorf("option with id %d does not exist", input.ID)
	}

	// Update the option in the database
	_, err = database.GetDBConnection().Exec("UPDATE options SET body = ?, correct = ? WHERE id = ?", input.Body, input.Correct, input.ID)
	if err != nil {
		return nil, err
	}

	return getOptionByID(input.ID)
}

// DeleteOption deletes an existing option.
// It validates the input, checks if the option exists, and deletes the option from the database.
func DeleteOption(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Extract the ID from the input arguments
	id, ok := args["id"].(int)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'id' parameter")
	}

	// Get the user ID from the context
	var userID = security.GetUserID(ctx)

	// Check if the option exists and if the user has access to it
	_, err := getQuestionByOptionID(id, *userID)
	if err != nil {
		return nil, fmt.Errorf("option with id %d does not exist", id)
	}

	// Delete the option from the database
	result, err := database.GetDBConnection().Exec("DELETE FROM options WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	// Check if the deletion was successful
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("option with id %d does not exist", id)
	}

	return true, nil
}

// getOptionByID retrieves an option by its ID from the database.
func getOptionByID(id int) (types.Option, error) {
	var option types.Option
	if err := database.GetDBConnection().Get(&option, "SELECT * FROM options WHERE id = ?", id); err != nil {
		return types.Option{}, err
	}
	return option, nil
}

// getQuestionByOptionID retrieves a question associated with an option by the option ID and user ID from the database.
func getQuestionByOptionID(optionID, userID int) (types.Question, error) {
	var question types.Question
	if err := database.GetDBConnection().Get(&question, `
		SELECT q.* 
			FROM options o
			JOIN questions q ON q.id = o.question_id
		WHERE o.id = ? AND q.user_id = ?`, optionID, userID); err != nil {
		return types.Question{}, err
	}
	return question, nil
}
