package repository

import (
	"github.com/bagasjs/go-backend-template/app/core"
	"github.com/bagasjs/go-backend-template/app/entity"
)

type UserRepository interface {
    ListAll() ([]entity.User, *core.Error)

    // Core Repository Methods
    Query(qb *core.QueryBuilder)([]entity.User, *core.Error)
    Insert(user entity.User) *core.Error 
    Update(id int, user entity.User) *core.Error 
    Destroy(id int) *core.Error 
}
