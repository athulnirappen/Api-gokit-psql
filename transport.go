package main

import (
    "context"
    "encoding/json"
    "net/http"
)

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req AddUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
