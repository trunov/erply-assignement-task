package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
)

type storage struct {
	db *sqlx.DB
}

func New(dsn string) (*storage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) GetCustomer(ctx context.Context, id int) (*erply.Customer, error) {
	customer := &erply.Customer{}
	query := `SELECT * FROM customers WHERE ID = $1`

	row := s.db.QueryRowxContext(ctx, query, id)
	if err := CustomScanCustomer(row, customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *storage) StoreCustomer(ctx context.Context, customer *erply.Customer) error {
	// Create a map to hold the parameter values
	params := make(map[string]interface{})

	// Use reflection to iterate over struct fields and populate the params map
	customerValue := reflect.ValueOf(customer).Elem()
	customerType := customerValue.Type()
	for i := 0; i < customerType.NumField(); i++ {
		field := customerType.Field(i)
		fieldName := strings.ToLower(field.Name)
		fieldValue := customerValue.Field(i).Interface()
		params[fieldName] = fieldValue
	}

	// Convert externalIDs to JSON-encoded string
	externalIDsJSON, err := json.Marshal(customer.ExternalIDs)
	if err != nil {
		return err
	}

	// Add externalIDs as a parameter
	params["externalids"] = string(externalIDsJSON)

	// Generate the query with placeholders
	var buf bytes.Buffer
	var keys []string
	var placeholders []string
	for key := range params {
		keys = append(keys, key)
		placeholders = append(placeholders, ":"+key)
	}
	fmt.Fprintf(&buf, "INSERT INTO customers (%s) VALUES (%s)", strings.Join(keys, ", "), strings.Join(placeholders, ", "))

	_, err = s.db.NamedExecContext(ctx, buf.String(), params)
	if err != nil {
		return err
	}

	return nil
}
