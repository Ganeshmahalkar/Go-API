package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockcollection struct {
}

func (*mockcollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func (*mockcollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

func TestInsert(t *testing.T) {
	mockCol := &mockcollection{}
	res, err := insertData(mockCol, User{"Rohit", 35})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}

func TestDelete(t *testing.T) {
	mockCol := &mockcollection{}
	res, err := deleteData(mockCol, User{"Akash", 30})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.DeleteResult{}, res)
}
