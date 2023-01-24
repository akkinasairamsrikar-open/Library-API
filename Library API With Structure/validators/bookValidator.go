package validators

import (
	"fmt"

	"github.com/Library-RestAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/govalidator"
)

// create a book validator and take input and bind json and validate the input and return bool and string

func BookValidatorHandler(c *gin.Context) (models.Book, error) {

	var bookObj models.Book
	// bind the temp to the request body
	if err := c.BindJSON(&bookObj); err != nil {
		return bookObj, errors.Wrap(err, "[BookValidatorHandler]: Errror binding to Json")
	}
	// create a map of rules
	rules := govalidator.MapData{
		"id":       []string{"required"},
		"title":    []string{"required", "min:3", "max:100"},
		"author":   []string{"required", "min:3", "max:100"},
		"quantity": []string{"required", "numeric", "min:1", "max:20"},
	}
	// create a map of messages
	messages := govalidator.MapData{
		"id":       []string{"required:ID is required"},
		"title":    []string{"required:Title is required", "min:Title should be at least 3 characters long", "max:Title should be at most 100 characters long"},
		"author":   []string{"required:Author is required", "min:Author should be at least 3 characters long", "max:Author should be at most 100 characters long"},
		"quantity": []string{"required:Quantity is required", "numeric:Quantity should be a number", "min:Quantity should be at least 1", "max:Quantity should be at most 20"},
	}
	// create a map of options to temp variable
	opts := govalidator.Options{
		Rules:    rules,
		Messages: messages,
		Data:     &bookObj,
	}
	// validate the input
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	// print the error
	fmt.Println(e)
	fmt.Println(bookObj)
	// if there is an error return false and the error message
	if len(e) > 0 {
		return bookObj, errors.Wrap(errors.New("Validator error"), "[BookValidatorHandler]: Errror binding to Json")
	}
	return bookObj, nil
}
