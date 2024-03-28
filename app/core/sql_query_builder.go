package core

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
    tableName string
    limit uint 
    columns []string
    conditions []string
    conditionValues []interface{}

}

func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{}
}

func (q *QueryBuilder) Table(name string) *QueryBuilder {
    q.tableName = name
    return q
}

func (q *QueryBuilder) Where(column string, operator string, value interface{}) *QueryBuilder {
    q.conditions = append(q.conditions, fmt.Sprintf("%s %s ?", column, operator))
    q.conditionValues = append(q.conditionValues, value)
    return q
}

func (q *QueryBuilder) Limit(amount uint) *QueryBuilder {
    q.limit = amount
    return q
}

func (q *QueryBuilder) Values() []interface{} {
    return q.conditionValues
}

func (q *QueryBuilder) ToString() string {
    var columnsString = "*"
    var conditionsString = ""
    if len(q.columns) != 0 {
        columnsString = strings.Join(q.columns, ",")
    }

    if len(q.conditions) != 0 {
        conditionsString = fmt.Sprintf("WHERE %s", strings.Join(q.conditions, ","))
    }

    if q.limit != 0 {
        return fmt.Sprintf("SELECT %s FROM %s %s LIMIT %d", columnsString, q.tableName, conditionsString, q.limit)
    } else {
        return fmt.Sprintf("SELECT %s FROM %s %s", columnsString, q.tableName, conditionsString)
    }
}

