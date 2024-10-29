package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http" 
	_ "github.com/lib/pq"
)

// connectDB initializes the database connection
func connectDB() (*sql.DB, error) {
    connStr := "user=postgres dbname=gokit password=athul sslmode=disable" 
    return sql.Open("postgres", connStr)
}

func main() {
    db, err := connectDB()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    svc := NewUserService(db)

    addUserHandler := httptransport.NewServer(
        makeAddUserEndpoint(svc),
        decodeAddUserRequest,
        encodeResponse,
    )
    getUserHandler := httptransport.NewServer(
        makeGetUserEndpoint(svc),
        decodeGetUserRequest,
        encodeResponse,
    )

    http.Handle("/add-user", addUserHandler)
    http.Handle("/get-user", getUserHandler)
    fmt.Println("server is running in the port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
