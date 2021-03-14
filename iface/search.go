package iface

/*
 * interface for search service, main entry
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

type ISearch interface {
	Quit()
	DocRemove(tag, docId string) bool
	DocSync(tag, docId string, jsonByte []byte) bool
	GetSuggest() ISuggest
	GetAgg() IAgg
	GetQuery() IQuery
	GetDoc() IDoc
	GetIndex(tag string) IIndex
	AddIndex(dir, tag string) bool
	AddNode(addr string) bool
	GetManager() IManager
}

