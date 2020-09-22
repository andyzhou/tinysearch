package iface

import "github.com/andyzhou/tinySearch/json"

/*
 * interface for doc
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type IDoc interface {
	GetCount(index IIndex) int64
	RemoveDoc(index IIndex, docId string) bool
	AddDoc(index IIndex, obj *json.DocJson) bool
}