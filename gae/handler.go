package gae

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/ebisu-voice-production/go-google-actions-api/api"
)

type GaeHandler struct {
	// UnmarshalConversationToken([]byte) (*ConversationToken, error)
	UnmarshalConversationToken interface{}
	// MarshalConversationToken(*ConversationToken) ([]byte, error)
	MarshalConversationToken interface{}
	// UnmarshalUserStorage([]byte) (*UserStorage, error)
	UnmarshalUserStorage interface{}
	// MarshalUserStorage(*UserStorage) ([]byte, error)
	MarshalUserStorage interface{}
	// HandleApp(ctx context.Context, req *api.AppRequest, res *api.AppResponse, token *ConversationToken, storage *UserStorage)
	HandleApp interface{}
}

func (g *GaeHandler) getConversationToken(ctx context.Context, req *api.AppRequest, t reflect.Type) reflect.Value {
	if g.UnmarshalConversationToken == nil {
		return reflect.Zero(t)
	}
	f := reflect.ValueOf(g.UnmarshalConversationToken)
	arg0 := reflect.ValueOf([]byte(req.GetConversationToken()))
	vs := f.Call([]reflect.Value{arg0})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to get conversationToken: %v", vs[1])
	}
	return vs[0]
}

func (g *GaeHandler) setConversationToken(ctx context.Context, req *api.AppRequest, res *api.AppResponse, v reflect.Value) {
	if g.MarshalConversationToken == nil {
		return
	}
	f := reflect.ValueOf(g.MarshalConversationToken)
	vs := f.Call([]reflect.Value{v})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to set conversationToken: %v", vs[1])
		return
	}
	res.ConversationToken = string(vs[0].Bytes())
}

func (g *GaeHandler) getUserStorage(ctx context.Context, req *api.AppRequest, t reflect.Type) reflect.Value {
	if g.UnmarshalUserStorage == nil {
		return reflect.Zero(t)
	}
	f := reflect.ValueOf(g.UnmarshalUserStorage)
	arg0 := reflect.ValueOf([]byte(req.GetUserStorage()))
	vs := f.Call([]reflect.Value{arg0})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to get conversationToken: %v", vs[1])
	}
	return vs[0]
}

func (g *GaeHandler) setUserStorage(ctx context.Context, req *api.AppRequest, res *api.AppResponse, v reflect.Value) {
	if g.MarshalUserStorage == nil {
		return
	}
	f := reflect.ValueOf(g.MarshalUserStorage)
	vs := f.Call([]reflect.Value{v})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to set conversationToken: %v", vs[1])
		return
	}
	res.UserStorage = string(vs[0].Bytes())
}

func (g *GaeHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf(ctx, "request body: %s", body)
	var req api.AppRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var res api.AppResponse
	f := reflect.ValueOf(g.HandleApp)
	tokenType := f.Type().In(3)
	storageType := f.Type().In(4)
	token := g.getConversationToken(ctx, &req, tokenType)
	storage := g.getUserStorage(ctx, &req, storageType)
	arg0 := reflect.ValueOf(ctx)
	arg1 := reflect.ValueOf(&req)
	arg2 := reflect.ValueOf(&res)
	f.Call([]reflect.Value{arg0, arg1, arg2, token, storage})
	g.setConversationToken(ctx, &req, &res, token)
	g.setUserStorage(ctx, &req, &res, storage)
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf(ctx, "response body: %s", js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
