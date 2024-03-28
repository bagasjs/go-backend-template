package application

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bagasjs/go-backend-template/app/repository"
	"github.com/bagasjs/go-backend-template/app/service"
)

type Application struct {
    UserRepository repository.UserRepository
    UserService *service.UserService

    // Private members
    db *sql.DB
}

type ApplicationConfig struct {

}

func (app *Application) Init(config ApplicationConfig) error {
    db, err := sql.Open("sqlite3", "./res/db.sqlite3")
    if err != nil {
        return err
    }

    app.db = db

    app.UserRepository = repository.NewUserRepository(app.db)
    app.UserService = service.NewUserService(app.UserRepository)

    return nil
}

func (app *Application) Destroy() {
    app.db.Close()
}

