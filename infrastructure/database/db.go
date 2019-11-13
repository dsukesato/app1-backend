package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
	user           = os.Getenv("CLOUDSQL_USER")
	password       = os.Getenv("CLOUDSQL_PASSWORD")
	database       = os.Getenv("CLOUDSQL_DATABASE")
)

const driverName = "mysql"

func init() {
	conn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", user, password, connectionName, database)

	db, err := sql.Open(driverName, conn)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db != nil {
			err := db.Close()
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("DB connection success!!")
			}
		}
	}()
	_, err = db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
}

/*type postsRepository struct{}

func (pr postsRepository) GetPosts() ([]*model.Post, error) {
	panic("implement me")
}

func NewPostsRepository() repository.PostsRepository {
	return &postsRepository{}
}*/
