package face

import (
	"errors"
	"github.com/andyzhou/tinySearch/iface"
	"sync"
)

/*
 * face for inter manager
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Manager struct {
	dataPath string
	indexes *sync.Map
	clients *sync.Map
}

//construct
func NewManager(dataPath string) *Manager{
	//self init
	this := &Manager{
		dataPath:dataPath,
		indexes:new(sync.Map),
		clients:new(sync.Map),
	}
	return this
}

//quit
func (f *Manager) Quit() {
	if f.clients == nil {
		return
	}
	sf := func(_, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		client.Quit()
		return true
	}
	f.clients.Range(sf)
}

//batch doc remove
func (f *Manager) DocsRemove(
					tag string,
					docIds []string,
				) bool {
	//basic check
	if tag == "" || docIds == nil {
		return false
	}
	if f.clients == nil {
		return false
	}

	//do doc sync on all clients
	sf := func(k, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		client.DocRemove(tag, docIds)
		return true
	}
	f.clients.Range(sf)
	return true
}

//doc remove from all clients
func (f *Manager) DocRemove(
					tag string,
					docId string,
				) error {
	//basic check
	if tag == "" || docId == "" {
		return errors.New("invalid parameter")
	}

	if f.clients == nil {
		return errors.New("no any client")
	}

	//do doc sync on all clients
	sf := func(k, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		client.DocRemove(tag, []string{docId})
		return true
	}
	f.clients.Range(sf)

	return nil
}

//doc sync to all clients
func (f *Manager) DocSync(
					tag string,
					docId string,
					jsonByte []byte,
				) error {
	//basic check
	if tag == "" || docId == "" || jsonByte == nil {
		return errors.New("invalid parameter")
	}

	if f.clients == nil {
		return errors.New("no any client")
	}

	//do doc sync on all clients
	sf := func(k, v interface{}) bool {
		client, ok := v.(*Client)
		if !ok {
			return false
		}
		client.DocSync(tag, docId, jsonByte)
		return true
	}
	f.clients.Range(sf)

	return nil
}

//remove client node
func (f *Manager) RemoveNode(
					addr string,
				) bool {
	//basic check
	if addr == "" || f.clients == nil {
		return false
	}

	//remove
	f.clients.Delete(addr)

	return true
}

//add client node
func (f *Manager) AddNode(
					addrArr ...string,
				) bool {
	//basic check
	if addrArr == nil || len(addrArr) <= 0 || f.clients == nil {
		return false
	}

	for _, addr := range addrArr {
		//check record
		_, ok := f.clients.Load(addr)
		if ok {
			continue
		}

		//init new client
		client := NewClient(addr)

		//sync into map
		f.clients.Store(addr, client)
	}

	return true
}

//get rand client
func (f *Manager) GetClient() iface.IClient {
	var (
		dstClient iface.IClient
		hasFound bool
	)
	if f.clients == nil {
		return nil
	}
	sf := func(key, val interface{}) bool {
		if hasFound {
			return false
		}
		//check client
		client, ok := val.(*Client)
		if !ok || client == nil {
			return false
		}
		if !client.isActive {
			return false
		}
		dstClient = client
		hasFound = true
		return true
	}
	f.clients.Range(sf)
	return dstClient
}

////////////////
//api for index
////////////////

//remove index
func (f *Manager) RemoveIndex(
					tag string,
				) bool {
	//basic check
	if tag == "" || f.indexes == nil {
		return false
	}

	//remove index
	f.indexes.Delete(tag)

	return true
}

//get search index
func (f *Manager) GetIndex(
					tag string,
				) iface.IIndex {
	//basic check
	if tag == "" || f.indexes == nil {
		return nil
	}

	//load record
	v, ok := f.indexes.Load(tag)
	if !ok {
		return nil
	}
	index, ok := v.(iface.IIndex)
	if !ok {
		return nil
	}

	return index
}

//add search index
func (f *Manager) AddIndex(
					tag string,
				) bool {
	//basic check
	if tag == "" || f.indexes == nil {
		return false
	}

	//check record
	_, ok := f.indexes.Load(tag)
	if ok {
		return false
	}

	//init new index
	index := NewIndex(f.dataPath, tag)

	//create index
	index.CreateIndex()

	//sync into map
	f.indexes.Store(tag, index)

	return true
}