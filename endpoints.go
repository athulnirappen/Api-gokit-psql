package main

import (
	"context"


    "github.com/go-kit/kit/endpoint"
)

type AddUserRequest struct {
	Name string `json:"name"`
}

type AddUserResponse struct {
    Msg string `json:"msg"`
	ID  int    `json:"id"`
	Err string `json:"err,omitempty"`
}

func makeAddUserEndpoint(svc UserService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddUserRequest)
		id, err := svc.AddUser(ctx, req.Name)
        if err != nil {
            return AddUserResponse{Err: err.Error()}, nil
        }
        return AddUserResponse{ID: id,Msg: "user is added"}, nil
    }
	
}

type GetUserRequest struct {
    ID int `json:"id"`
}

type GetUserResponse struct {
    User User   `json:"user"`
    Err  string `json:"err,omitempty"`
}

func makeGetUserEndpoint(svc UserService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetUserRequest)
        user, err := svc.GetUser(ctx, req.ID)
        if err != nil {
            return GetUserResponse{Err: err.Error()}, nil
        }
        return GetUserResponse{User: user}, nil
    }
}