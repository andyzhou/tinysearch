package iface

import (
	"github.com/andyzhou/tinySearch/json"
	"github.com/blevesearch/bleve/v2"
)

/*
 * interface for query
 */

type IQuery interface {
	QueryAll(index IIndex) (*json.SearchResultJson, error)
	Query(index IIndex, json *json.QueryOptJson) (*json.SearchResultJson, error)
	BuildSearchReq(json *json.QueryOptJson) *bleve.SearchRequest
}