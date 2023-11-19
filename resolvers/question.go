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

	for i := range questions {
		if err := database.GetDBConnection().Select(&questions[i].Options, "SELECT * FROM options WHERE question_id = ?", questions[i].ID); err != nil {
			return nil, err
		}
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
	if err := database.GetDBConnection().Get(&question, "SELECT * FROM questions WHERE id = ? AND user_id = ?", id, userID); err != nil {
		return types.Question{}, err
	}

	// var options []types.Option
	if err := database.GetDBConnection().Select(&question.Options, "SELECT * FROM options WHERE question_id = ?", question.ID); err != nil {
		return types.Question{}, err
	}
	return question, nil
}

func AddQuestion(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	var input struct {
		Body    string
		Options []struct {
			Body    string
			Correct bool
		}
	}

	mapstructure.Decode(args, &input)

	var userID = security.GetUserID(ctx)

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
