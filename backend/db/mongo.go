package mongo 

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Backgroud(), "Insert MOngo URL", nil)