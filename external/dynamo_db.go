package external

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/guregu/dynamo/v2"
	"github.com/guregu/dynamo/v2/dynamodbiface"
)

//go:generate mockgen -source=dynamo_db.go -destination=mock/dynamo_db.go
type DynamoDatabase interface {
	Client() dynamodbiface.DynamoDBAPI
	CreateTable(name string, from interface{}) *dynamo.CreateTable
	GetTx() *dynamo.GetTx
	ListTables() *dynamo.ListTables
	Table(name string) dynamo.Table
	WriteTx() *dynamo.WriteTx
}

func GetDynamoDatabase(cfg aws.Config) DynamoDatabase {
	return dynamo.New(cfg)
}

type dynamoAdapter struct {
	db    DynamoDatabase
	table *string
}

type DynamoAdapter interface {
	SetTable(table string)
	GetAll() (value interface{}, err error)
	GetOneByKey(key string, valueKey interface{}) (value interface{}, err error)
	Create(value interface{}) (err error)
	UpdateValue(key string, valueKey interface{}, keyToUpdate string, valueToUpdate interface{}) (updatedValue interface{}, err error)
}

func NewDynamoAdapter(db DynamoDatabase) DynamoAdapter {
	return &dynamoAdapter{
		db:    db,
		table: nil,
	}
}

func (d *dynamoAdapter) SetTable(table string) {
	d.table = &table
}

// GetAll implements DynamoAdapter.
func (d *dynamoAdapter) GetAll() (value interface{}, err error) {
	err = d.db.Table(*d.table).Scan().All(context.TODO(), &value)
	return
}

func (d *dynamoAdapter) GetOneByKey(key string, valueKey interface{}) (value interface{}, err error) {
	err = d.db.Table(*d.table).Get(key, valueKey).One(context.TODO(), &value)
	return
}

func (d *dynamoAdapter) Create(value interface{}) (err error) {
	err = d.db.Table(*d.table).Put(value).Run(context.TODO())
	return
}

func (d *dynamoAdapter) UpdateValue(key string, valueKey interface{}, keyToUpdate string, valueToUpdate interface{}) (updatedValue interface{}, err error) {
	err = d.db.Table(*d.table).Update(key, valueKey).
		Set(keyToUpdate, valueToUpdate).
		Value(context.TODO(), &updatedValue)
	return
}
