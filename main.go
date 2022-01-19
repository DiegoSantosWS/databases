package main

import (
	"context"
	"time"

	"github.com/DiegoSantosWS/databases/elasticsearch/es01db"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*2)
	defer cancel()
	ctx = es01db.GetConnect(ctx)
	// elasticsearch.Insert(ctx)
	// elasticsearch.Update(ctx)
}
