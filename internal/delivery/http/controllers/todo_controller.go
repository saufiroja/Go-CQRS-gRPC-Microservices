package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/handlers"
	"net/http"
)

type Controllers struct {
	handler handlers.TodoHandler
}

func NewControllers(handler handlers.TodoHandler) ITodoController {
	return &Controllers{
		handler: handler,
	}
}

func (c *Controllers) InsertTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Command.InsertTodoCommand.Handle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) GetAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Query.GetAllTodoQuery.Handle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) GetTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Query.GetTodoByIdQuery.Handle(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) UpdateTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Command.UpdateTodoCommand.Handle(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) UpdateStatusTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Command.UpdateStatusTodoByIdCommand.Handle(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) DeleteTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.handler.Command.DeleteTodoByIdCommand.Handle(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}