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

func GetAllQuestions(ctx context.Context) (interface{}, error) {
	var userID = security.GetUserID(ctx)

	var questions []types.Question
	err := database.GetDBConnection().Select(&questions, "SELECT * FROM questions WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func GetQuestionByID(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	id, ok := args["id"].(int)
	if !ok {
		return nil, errors.New("could not parse id")
	}

	var userID = security.GetUserID(ctx)

	return getQuestionByID(id, *userID)
}

func getQuestionByID(id, userID int) (types.Question, error) {
	var question types.Question
	err := database.GetDBConnection().Get(&question, "SELECT * FROM questions WHERE id= ? AND user_id", id, userID)
	if err != nil {
		return types.Question{}, err
	}
	return question, err
}

func GetOptionsByQuestionID(questionID int) (interface{}, error) {
	var options []types.Option
	err := database.GetDBConnection().Select(&options, "SELECT * FROM options WHERE question_id = ? ORDER BY id", questionID)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func AddQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	body, ok := args["body"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'body' parameter")
	}

	var userID = security.GetUserID(ctx)

	// Insert the new question into the database
	result, err := database.GetDBConnection().Exec("INSERT INTO questions (body, user_id) VALUES (?, ?)", body, userID)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted question
	id, _ := result.LastInsertId()

	// Return the newly added question
	return getQuestionByID(int(id), *userID)
}

func UpdateQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	var input struct {
		ID   int
		Body string
	}
	if err := mapstructure.Decode(args, &input); err != nil {
		return nil, err
	}

	var userID = security.GetUserID(ctx)

	// Update the question in the database
	_, err := database.GetDBConnection().Exec("UPDATE questions SET body = ? WHERE id = ? AND user_id = ?", input.Body, input.ID, userID)
	if err != nil {
		return nil, err
	}

	// Return the updated question
	return getQuestionByID(input.ID, *userID)
}

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
