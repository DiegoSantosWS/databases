package es01db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type Client struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `josn:"client_address"`
	Phone   Phone   `json:"phone"`
}

type Address struct {
	Address string `json:"address"`
	Number  string `json:"number"`
}

type Phone struct {
	NumType string `json:"num_type"`
	Number  string `json:"number"`
}

// LoadObjectID ...
func LoadObjectID(ctx context.Context, nameIdx string, obj interface{}) {
	client := ctx.Value(ClientKey).(*elasticsearch.Client)
	objs := obj.(Client)
	res, err := client.Index(nameIdx, esutil.NewJSONReader(obj),
		client.Index.WithDocumentID(objs.ID))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// Update ...
func Update(ctx context.Context, nameIdx, id string, obj interface{}) {
	client := ctx.Value(ClientKey).(*elasticsearch.Client)
	bdy, err := json.Marshal(obj)
	if err != nil {
		return
	}

	req := esapi.UpdateRequest{
		Index:      nameIdx,
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	res, err := req.Do(ctx, client)
	if err != nil {
		log.Println("do:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return
	}

	log.Println(res)
}

// Update ...
func Insert(ctx context.Context, nameIdx, id string, obj interface{}) {
	client := ctx.Value(ClientKey).(*elasticsearch.Client)
	bdy, err := json.Marshal(obj)
	if err != nil {
		return
	}

	req := esapi.CreateRequest{
		Index:      nameIdx,
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	res, err := req.Do(ctx, client)
	if err != nil {
		log.Println("do:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return
	}

	log.Println(res)
}
