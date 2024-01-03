package internals

import (
	"fmt"
	"strings"
)

type Querybuilder struct {
	Type                  string
	Table                 string
	Conditions            string
	IsMethodAlreadyCalled bool
	Join                  string
	Error                 error
	Fields                []string
}

func NewQuery() *Querybuilder {
	return &Querybuilder{}
}

func (query *Querybuilder) FROMTABLE(targetedTable string) *Querybuilder {
	query.Table = targetedTable
	return query
}

func (query *Querybuilder) SELECT(fields ...string) *Querybuilder {
	if query.IsMethodAlreadyCalled {
		return &Querybuilder{Error: fmt.Errorf("A query method has been already invoked")}
	}
	query.IsMethodAlreadyCalled = true
	query.Type = "SELECT"
	query.Fields = fields
	return query
}
func (query *Querybuilder) JOIN(joinType, joinTable, onCondition string) *Querybuilder {
	if query.Type != "SELECT" {
		return &Querybuilder{Error: fmt.Errorf("JOIN statement is only supported for SELECT queries")}
	}
	query.Join += fmt.Sprintf(" %s %s ON %s", joinType, joinTable, onCondition)
	return query
}
func (query *Querybuilder) WHERE(condition string) *Querybuilder {
	query.Conditions = condition
	return query
}

func (query *Querybuilder) DELETE(fields ...string) *Querybuilder {
	if query.IsMethodAlreadyCalled {
		return &Querybuilder{Error: fmt.Errorf("A query method has been already invoked")}
	}
	query.IsMethodAlreadyCalled = true
	query.Type = "DELETE"
	query.Fields = fields
	return query
}

func (query *Querybuilder) INSERT(fields ...string) *Querybuilder {
	if query.IsMethodAlreadyCalled {
		return &Querybuilder{Error: fmt.Errorf("A query method has been already invoked")}
	}
	query.IsMethodAlreadyCalled = true
	query.Type = "INSERT INTO"
	query.Fields = fields
	return query
}

func (query *Querybuilder) Build() (string, error) {
	if query.Error != nil {
		return "", query.Error
	}
	if query.Type == "" || query.Table == "" {
		return "", fmt.Errorf("The query type and table must be specified")
	}
	switch query.Type {
	case "SELECT":
		queryString := fmt.Sprintf("SELECT %s FROM %s", strings.Join(query.Fields, ","), query.Table)
		if query.Join != "" {
			queryString += query.Join
		}
		if query.Conditions != "" {
			queryString += fmt.Sprintf(" WHERE %s", query.Conditions)
		}
		return queryString, nil
	case "DELETE":
		if query.Conditions == "" {
			return "", fmt.Errorf("DELETE query must have conditions")
		}
		return fmt.Sprintf("DELETE FROM %s WHERE %s", query.Table, query.Conditions), nil
	case "INSERT INTO":
		if len(query.Fields) == 0 {
			return "", fmt.Errorf("Fields for INSERT must be specified")
		}
		valuePlaceholders := strings.Repeat("?, ", len(query.Fields)-1) + "?"
		return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", query.Table, strings.Join(query.Fields, ","), valuePlaceholders), nil
	default:
		return "", fmt.Errorf("Unsupported query type: %s", query.Type)
	}
}
