package core

import (
	"fmt"
	"strings"
)

type UpdateBuilder struct {
    tableName string

    columns []string
    columnValues []interface{}
    conditions []string
    conditionValues []interface{}
}

func NewUpdateBuilder() *UpdateBuilder {
    return &UpdateBuilder{}
}

func (b *UpdateBuilder) Table(name string) *UpdateBuilder {
    b.tableName = name
    return b
}

func (b *UpdateBuilder) Where(column string, operator string, value interface{}) *UpdateBuilder {
    b.conditions = append(b.conditions, fmt.Sprintf("%s %s ?", column, operator))
    b.conditionValues = append(b.conditionValues, value)
    return b
}

func (b *UpdateBuilder) Column(name string, value interface{}) *UpdateBuilder {
    b.columns = append(b.columns, fmt.Sprintf("%s=?", name))
    b.columnValues = append(b.columnValues, value) 
    return b
}


func (b *UpdateBuilder) Values() []interface{} {
    values := make([]interface{}, 0)
    values = append(values, b.columnValues...)
    values = append(values, b.conditionValues...)
    return values
}

func (b *UpdateBuilder) ToString() string {
    sets := strings.Join(b.columns, ",")
    if len(b.conditions) > 0 {
        conditions := strings.Join(b.conditions, ",")
        return fmt.Sprintf("UPDATE %s SET %s WHERE %s", b.tableName, sets, conditions)
    } else {
        return fmt.Sprintf("UPDATE %s SET %s", b.tableName, sets)
    }
}
