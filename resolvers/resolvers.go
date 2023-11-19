package resolvers

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/database"
	"github.com/togglhire/backend-homework/types"
)

func GetAllQuestions() (interface{}, error) {
	var questions []types.Question
	err := database.GetDBConnection().Select(&questions, "SELECT * FROM questions")
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func GetQuestionByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("could not parse id")
	}

	return getQuestionByID(id)
}

func getQuestionByID(id int) (types.Question, error) {
	var question types.Question
	err := database.GetDBConnection().Get(&question, "SELECT * FROM questions WHERE id=? ORDER BY id", id)
	if err != nil {
		return types.Question{}, err
	}
	return question, err
}

func GetOptionsByQuestionID(questionID int) (interface{}, error) {
	var options []types.Option
	err := database.GetDBConnection().Select(&options, "SELECT * FROM options WHERE question_id=? ORDER BY id", questionID)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func AddQuestion(p graphql.ResolveParams) (interface{}, error) {
	body, ok := p.Args["body"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'body' parameter")
	}

	// Insert the new question into the database
	result, err := database.GetDBConnection().Exec("INSERT INTO questions (body) VALUES (?)", body)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted question
	id, _ := result.LastInsertId()

	// Return the newly added question
	return getQuestionByID(int(id))
}

func UpdateQuestion(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'id' parameter")
	}

	body, ok := p.Args["body"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'body' parameter")
	}

	// Update the question in the database
	_, err := database.GetDBConnection().Exec("UPDATE questions SET body = ? WHERE id = ?", body, id)
	if err != nil {
		return nil, err
	}

	// Return the updated question
	return getQuestionByID(id)
}

func DeleteQuestion(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'id' parameter")
	}

	// Update the question in the database
	result, err := database.GetDBConnection().Exec("DELETE FROM questions WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, fmt.Errorf("question with id %d does not exist", id)
	}

	return true, nil
}
