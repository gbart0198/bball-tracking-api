package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gbart0198/bball-tracker-api/types"
	_ "github.com/microsoft/go-mssqldb"
)

type MSSQLStorage struct {
	connStr string
	db    *sql.DB
}

func NewMSSQLStorage(connStr string) *MSSQLStorage {
	return &MSSQLStorage{
		connStr: connStr,
	}
}

func (store *MSSQLStorage) Connect() {
	db, err := sql.Open("sqlserver", store.connStr)
	fmt.Println("Database connected")

	if err != nil {
		log.Fatal(err)
	}
	store.db = db
}

func (store *MSSQLStorage) GetUser(userID string) *types.User {
    fmt.Println("Looking for users with user id: ", userID)
    user := &types.User{}
    query := `SELECT TOP 1 UserId, FirstName, LastName, Email from dbo.Users WHERE UserId = @userID`
    err := store.db.QueryRow(query, sql.Named("userID", userID)).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)

    if err != nil {
        log.Fatal(err)
    }

    return user
}
