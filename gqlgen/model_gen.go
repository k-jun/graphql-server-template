// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

type TodoCreateInput struct {
	UserID  string `json:"userID"`
	Message string `json:"message"`
}

type TodoDeleteInput struct {
	ID string `json:"id"`
}

type TodoUpdateInput struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Done    bool   `json:"done"`
}

type UserCreateInput struct {
	Name string `json:"name"`
}