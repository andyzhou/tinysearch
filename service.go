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
type Service struct {
	manager iface.IManager
	rpc iface.IRpc
}

//construct
func NewService(dataPath string, rpcPort int) *Service {
	//self init
	this := &Service{
		manager: face.NewManager(dataPath),
	}
	//init rpc
	this.rpc = face.NewRpc(rpcPort, this.manager)
	return this
}

//quit
func (f *Service) Quit() {
	f.manager.Quit()
	f.rpc.Stop()
}

//doc remove from batch node
func (f *Service) DocRemove(
					tag string,
					docId string,
				) error {
	return f.manager.DocRemove(tag, docId)
}

//doc sync into batch node
func (f *Service) DocSync(
					tag string,
					docId string,
					jsonByte []byte,
				) error {
	return f.manager.DocSync(tag, docId, jsonByte)
}

//get suggest face
func (f *Service) GetSuggest() iface.ISuggest {
	return f.manager.GetSuggest()
}

//get agg face
func (f *Service) GetAgg() iface.IAgg {
	return f.manager.GetAgg()
}

//get query face
func (f *Service) GetQuery() iface.IQuery {
	return f.manager.GetQuery()
}

//get doc face
func (f *Service) GetDoc() iface.IDoc {
	return f.manager.GetDoc()
}

//get index face
func (f *Service) GetIndex(
					tag string,
				) iface.IIndex {
	return f.manager.GetIndex(tag)
}

//add index
func (f *Service) AddIndex(
					tag string,
				) bool {
	return f.manager.AddIndex(tag)
}

//add rpc node
func (f *Service) AddNode(
					addr ...string,
				) bool {
	return f.manager.AddNode(addr...)
}

//get manager face
func (f *Service) GetManager() iface.IManager {
	return f.manager
}