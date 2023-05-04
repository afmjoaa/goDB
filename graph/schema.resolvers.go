package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"
	"goDB/db_operations"
	"goDB/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todos := &model.Todo{
		ID:   fmt.Sprint(len(r.TodoList) + 1),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("Joaa%s", fmt.Sprint(len(r.TodoList)+1)),
		},
	}

	r.TodoList = append(r.TodoList, todos)
	return todos, nil
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, input model.TableData) (string, error) {
	return db_operations.CreateTablePk(input.TableName, input.PrimaryKey), nil
}

// CreateTableRow is the resolver for the createTableRow field.
func (r *mutationResolver) CreateTableRow(ctx context.Context, input model.TableRowData) (string, error) {
	return db_operations.Create(), nil
}

// UpdateTable is the resolver for the updateTable field.
func (r *mutationResolver) UpdateTable(ctx context.Context, input model.UpdateTableName) (string, error) {
	return db_operations.UpdateTableName(input.OldTableName, input.NewTableName), nil
}

// DeleteTable is the resolver for the deleteTable field.
func (r *mutationResolver) DeleteTable(ctx context.Context, input model.DeleteTableData) (string, error) {
	return db_operations.DeleteTable(input.TableName), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodoList, nil
}

// ReadTable is the resolver for the readTable field.
func (r *queryResolver) ReadTable(ctx context.Context, input model.ReadAllTableData) (string, error) {
	_, jsonStr := db_operations.ReadAll(input.TableName)
	return jsonStr, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
