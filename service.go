package main

import (
	"context"
	"database/sql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	AddUser(ctx context.Context, name string) (int, error)
	GetUser(ctx context.Context, id int) (User, error)
}

type userService struct {
    db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
    return &userService{db: db}
}

func (s *userService) AddUser(ctx context.Context,name string) (int,error) {
	var id int
    err := s.db.QueryRowContext(ctx, "INSERT INTO users(name) VALUES($1) RETURNING id", name).Scan(&id)
    return id, err
}

func (s *userService) GetUser(ctx context.Context,id int) (User,error) {
	var user User
    err := s.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
    return user, err
}