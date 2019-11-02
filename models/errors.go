package models

import (
	"database/sql"
	"errors"
	"log"

	"github.com/evalphobia/go-timber/timber"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

/*Handlerr funcao encarregada para tratar exibir erros de banco*/
func Handlerr(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

/*NoRow retorna true se ocorreu algum erro ou nao existe a linha*/
func NoRow(err error) bool {
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return true
	}
	return false
}

/*HandleError funcao encarregada para tratar exibir erros de banco*/
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

/*HandleErrorTimber tratador geral de erros com log no timber.io*/
func HandleErrorTimber(err error, cli *timber.Client) {
	if err != nil {
		cli.Err(err.Error())
		log.Fatal(err)
	}
}
