package tinySearch

import (
	"github.com/andyzhou/tinySearch/face"
	"github.com/andyzhou/tinySearch/iface"
)

/*
 * service api
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Search struct {
	suggest iface.ISuggest
	agg iface.IAgg
	query iface.IQuery
	doc iface.IDoc
	manager iface.IManager
	rpc iface.IRpc
}

//construct
func NewSearch(rpcPort int) *Search {
	//self init
	this := &Search{
		manager: face.NewManager(),
		doc:     face.NewDoc(),
		query:   face.NewQuery(),
		agg:     face.NewAgg(),
		suggest: face.NewSuggest(),
	}
	//init rpc
	this.rpc = face.NewRpc(rpcPort, this.manager)
	return this
}

//quit
func (f *Search) Quit() {
	f.rpc.Stop()
}

//doc remove from batch node
func (f *Search) DocRemove(
					tag string,
					docId string,
				) error {
	return f.manager.DocRemove(tag, docId)
}

//doc sync into batch node
func (f *Search) DocSync(
					tag string,
					docId string,
					jsonByte []byte,
				) error {
	return f.manager.DocSync(tag, docId, jsonByte)
}

//get suggest face
func (f *Search) GetSuggest() iface.ISuggest {
	return f.suggest
}

//get agg face
func (f *Search) GetAgg() iface.IAgg {
	return f.agg
}

//get query face
func (f *Search) GetQuery() iface.IQuery {
	return f.query
}

//get doc face
func (f *Search) GetDoc() iface.IDoc {
	return f.doc
}

//get index face
func (f *Search) GetIndex(
					tag string,
				) iface.IIndex {
	return f.manager.GetIndex(tag)
}

//add index
func (f *Search) AddIndex(
					dir, tag string,
				) bool {
	return f.manager.AddIndex(dir, tag)
}

//add rpc node
func (f *Search) AddNode(
					addr ...string,
				) bool {
	return f.manager.AddNode(addr...)
}

//get manager face
func (f *Search) GetManager() iface.IManager {
	return f.manager
}