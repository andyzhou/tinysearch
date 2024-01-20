package tinysearch

import (
	"errors"
	"fmt"
	"github.com/andyzhou/tinysearch/iface"
	"github.com/andyzhou/tinysearch/json"
	"github.com/andyzhou/tinysearch/rpc"
	"log"
	"sync"
)

/*
 * client api
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 * - used for rpc mode
 */

//query opt kind
const (
	QueryOptKindOfGen = iota
	QueryOptKindOfAgg
	QueryOptKindOfSuggest
)

//others
const (
	SyncChanSize = 1024 * 5
	RemoveChanSize = 1024 * 5
)

//inter struct
type (
	syncDocReq struct {
		indexTag string
		docId string
		docJson []byte
	}

	removeDocReq struct {
		indexTag string
		docIds []string
	}
)

//face info
type Client struct {
	rpcClients map[string]iface.IRpcClient //address -> rpcClient
	syncChan chan syncDocReq //sync doc
	removeChan chan removeDocReq //remove batch docs
	closeChan chan struct{}
	sync.RWMutex
}

//construct
func NewClient() *Client {
	//self init
	self := &Client{
		rpcClients:make(map[string]iface.IRpcClient),
		syncChan: make(chan syncDocReq, SyncChanSize),
		removeChan:make(chan removeDocReq, RemoveChanSize),
		closeChan:make(chan struct{}, 1),
	}
	go self.runMainProcess()
	return self
}

//quit
func (f *Client) Quit() {
	var (
		m any = nil
	)
	defer func() {
		if err := recover(); err != m {
			log.Printf("tinysearch.Client:Quit panic, err:%v", err)
		}
	}()
	if f.rpcClients != nil {
		for _, client := range f.rpcClients {
			client.Quit()
		}
	}
	if f.closeChan != nil {
		close(f.closeChan)
	}
}

//suggest doc
func (f *Client) DocSuggest(
					indexTag string,
					optJson *json.QueryOptJson,
				) (*json.SuggestsJson, error) {
	//check
	if indexTag == "" || optJson == nil {
		return nil, errors.New("invalid parameter")
	}

	//get rpc client
	client := f.getClient()
	if client == nil {
		return nil, errors.New("can't get active rpc client")
	}

	optJsonByte, err := optJson.Encode()
	if err != nil {
		return nil, err
	}

	//call api
	jsonByte, err := client.DocQuery(
		QueryOptKindOfSuggest,
		indexTag,
		optJsonByte,
	)
	if err != nil {
		return nil, err
	}

	//analyze result
	if jsonByte == nil {
		return nil, nil
	}

	//format result
	resultJson := json.NewSuggestsJson()
	err = resultJson.Decode(jsonByte)
	if err != nil {
		return nil, err
	}
	return resultJson, nil
}

//agg doc
func (f *Client) DocAgg(
					indexTag string,
					optJson *json.QueryOptJson,
				) (*json.AggregatesJson, error) {
	//check
	if indexTag == "" || optJson == nil {
		return nil, errors.New("invalid parameter")
	}

	//get rpc client
	client := f.getClient()
	if client == nil {
		return nil, errors.New("can't get active rpc client")
	}

	optJsonByte, err := optJson.Encode()
	if err != nil {
		return nil, err
	}

	//call api
	jsonByte, err := client.DocQuery(
									QueryOptKindOfAgg,
									indexTag,
									optJsonByte,
								)
	if err != nil {
		return nil, err
	}

	//analyze result
	if jsonByte == nil {
		return nil, nil
	}

	//format result
	resultJson := json.NewAggregatesJson()
	err = resultJson.Decode(jsonByte)
	if err != nil {
		return nil, err
	}
	return resultJson, nil
}

//query doc
func (f *Client) DocQuery(
					indexTag string,
					optJson *json.QueryOptJson,
				) (*json.SearchResultJson, error) {
	//check
	if indexTag == "" || optJson == nil {
		return nil, errors.New("invalid parameter")
	}

	//get rpc client
	client := f.getClient()
	if client == nil {
		return nil, errors.New("can't get active rpc client")
	}

	optJsonByte, err := optJson.Encode()
	if err != nil {
		return nil, err
	}

	//call api
	jsonByte, err := client.DocQuery(
								QueryOptKindOfGen,
								indexTag,
								optJsonByte,
							)
	if err != nil {
		return nil, err
	}

	//analyze result
	if jsonByte == nil {
		return nil, nil
	}

	//format result
	resultJson := json.NewSearchResultJson()
	err = resultJson.Decode(jsonByte)
	if err != nil {
		return nil, err
	}
	return resultJson, nil
}

//remove doc
func (f *Client) DocRemove(
					indexTag string,
					docIds ...string,
				) error {
	var (
		m any = nil
	)
	//check
	if indexTag == "" || docIds == nil {
		return errors.New("invalid parameter")
	}
	if f.rpcClients == nil {
		return errors.New("no any active rpc client")
	}

	//defer
	defer func() {
		if e := recover(); e != m {
			log.Printf("client.DocRemove panic, err:%v\n", e)
		}
	}()

	//init request
	req := removeDocReq{
		indexTag: indexTag,
		docIds: docIds,
	}

	//send to chan
	select {
	case f.removeChan <- req:
	}
	return nil
}

//get doc
func (f *Client) DocGet(
					indexTag string,
					docIds ...string,
				) ([][]byte, error) {
	//check
	if indexTag == "" || docIds == nil {
		return nil, errors.New("invalid parameter")
	}

	//get rpc client
	client := f.getClient()
	if client == nil {
		return nil, errors.New("can't get active rpc client")
	}

	//call rpc api
	return client.DocGet(indexTag, docIds...)
}

//add sync
//used for add, sync doc, run on all nodes
func (f *Client) DocSync(
					indexTag, docId string,
					docJson []byte,
				) error {
	var (
		m any = nil
	)
	//check
	if indexTag == "" || docId == "" || docJson == nil {
		return errors.New("invalid parameter")
	}
	if f.rpcClients == nil {
		return errors.New("no any active rpc client")
	}

	//defer
	defer func() {
		if e := recover(); e != m {
			log.Printf("client.DocSync panic, err:%v\n", e)
			return
		}
	}()

	//init request
	req := syncDocReq{
		indexTag: indexTag,
		docId: docId,
		docJson: docJson,
	}

	//send to chan
	select {
	case f.syncChan <- req:
	}
	return nil
}

//create index
func (f *Client) CreateIndex(indexTag string) error {
	//check
	if indexTag == "" {
		return errors.New("invalid parameter")
	}
	//get rpc client
	client := f.getClient()
	if client == nil {
		return errors.New("can't get active rpc client")
	}
	//call rpc api
	err := client.IndexCreate(indexTag)
	return err
}

//add search service nodes
func (f *Client) AddNodes(nodes ... string) error {
	//check
	if nodes == nil || len(nodes) <= 0 {
		return errors.New("invalid parameter")
	}
	//check and init new rpc client
	for _, node := range nodes {
		//check
		v, ok := f.rpcClients[node]
		if ok && v != nil {
			continue
		}
		//create new rpc client
		rpcClient := rpc.NewRpcClient(node)
		f.rpcClients[node] = rpcClient
	}
	return nil
}

//////////////
//private func
//////////////

//run main process
func (f *Client) runMainProcess() {
	var (
		m any = nil
		syncReq syncDocReq
		removeReq removeDocReq
		isOk bool
	)

	//defer
	defer func() {
		if err := recover(); err != m {
			log.Println("Client:mainProcess panic, err:", err)
		}
		close(f.removeChan)
		close(f.syncChan)
		close(f.closeChan)
	}()

	//async loop
	for {
		select {
		case syncReq, isOk = <- f.syncChan:
			if isOk {
				//sync doc
				f.syncDoc(&syncReq)
			}
		case removeReq, isOk = <- f.removeChan:
			if isOk {
				//remove relate doc by ids
				f.removeBatchDocByIds(&removeReq)
			}

		case <- f.closeChan:
			return
		}
	}
}

//sync doc
func (f *Client) syncDoc(req *syncDocReq) bool {
	var (
		bRet bool
	)

	//check
	if req == nil || req.docJson == nil {
		return false
	}

	//run on all rpc clients
	succeed := 0
	failed := 0
	for _, client := range f.rpcClients {
		if !client.IsActive() {
			failed++
			continue
		}
		bRet = client.DocSync(req.indexTag, req.docId, req.docJson)
		if bRet {
			succeed++
		}else{
			failed++
		}
	}
	if failed > 0 {
		info := fmt.Sprintf("failed:%v, succeed:%v", failed, succeed)
		log.Printf("client:syncDoc, %v\n", info)
	}
	return true
}

//remove batch doc by ids
func (f *Client) removeBatchDocByIds(req *removeDocReq) bool {
	var (
		bRet bool
	)
	//check
	if req == nil || req.docIds == nil {
		return false
	}
	//run on all rpc clients
	succeed := 0
	failed := 0
	for _, client := range f.rpcClients {
		if !client.IsActive() {
			failed++
			continue
		}
		bRet = client.DocRemove(req.indexTag, req.docIds...)
		if bRet {
			succeed++
		}else{
			failed++
		}
	}
	if failed > 0 {
		info := fmt.Sprintf("failed:%v, succeed:%v", failed, succeed)
		log.Printf("client:removeBatchDocByIds, %v\n", info)
	}
	return true
}

//get rand active rpc client
func (f *Client) getClient() iface.IRpcClient {
	if f.rpcClients == nil {
		return nil
	}
	for _, client := range f.rpcClients {
		if client.IsActive() {
			return client
		}
	}
	return nil
}