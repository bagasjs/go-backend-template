package repository

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/bagasjs/go-backend-template/app/core"
	"github.com/bagasjs/go-backend-template/app/entity"
)


type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Query(query *core.QueryBuilder) ([]entity.User, *core.Error) {
    query.Table("internal__users")
    rows, err := repo.db.Query(query.ToString(), query.Values()...)
    if err != nil {
        return []entity.User{}, core.NewError(http.StatusInternalServerError, err.Error())
    }
    defer rows.Close()
    users := []entity.User{}
    for rows.Next() {
        user := entity.User{}
        err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Created, &user.Updated)
        if err != nil {
            return []entity.User{}, core.NewError(http.StatusNotImplemented, err.Error())
        }
        users = append(users, user)
    }

    return users, nil
}

func (repo *userRepositoryImplSQLite3) ListAll() ([]entity.User, *core.Error) {
    return repo.Query(core.NewQueryBuilder())
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) *core.Error {
    stmt, err := repo.db.Prepare("INSERT INTO `internal__users` (`name`, `email`, `password`, `created`, `updated`) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return core.NewError(http.StatusInternalServerError, err.Error())
    }
    defer stmt.Close()
    now := time.Now()
    user.Created = now.Format(time.RFC3339)
    user.Updated = user.Created
    _, err = stmt.Exec(user.Name, user.Email, user.Password, user.Created, user.Updated)
    if err != nil {
        return core.NewError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) Update(id int, userData entity.User) *core.Error {
    users, err := repo.Query(core.NewQueryBuilder().Limit(1).Where("id", "=", id))
    if err != nil {
        return err
    }

    now := time.Now()
    user := users[0]
    user.Name  = core.Choose(len(userData.Name) != 0, userData.Name, user.Name).(string)
    user.Email = core.Choose(len(userData.Email) != 0, userData.Email, user.Email).(string)
    user.Password = core.Choose(len(userData.Password) != 0, userData.Name, user.Password).(string)
    user.Updated = now.Format(time.RFC3339)

    stmt, gerr := repo.db.Prepare("UPDATE `internal__users` SET name=?, email=?, password=?, updated=? WHERE id=?");
    if gerr != nil {
        return core.NewError(http.StatusInternalServerError, gerr.Error())
    }
    _, gerr = stmt.Exec(user.Name, user.Email, user.Password, user.Updated, id)
    if gerr != nil {
        return core.NewError(http.StatusBadRequest, gerr.Error())
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) Destroy(id int) *core.Error {
    _, err := repo.db.Exec("DELETE FROM `internal__users` WHERE id=?", id)
    if err != nil {
        return core.NewError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func NewUserRepository(conn *sql.DB) UserRepository {
    return &userRepositoryImplSQLite3 {
        db: conn,
    };
}

