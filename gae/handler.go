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

type key int

const (
	KEY_UNKNOWN key = iota
	KEY_URL_PATH
)

type AppHandler struct {
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

func (a *AppHandler) getConversationToken(ctx context.Context, req *api.AppRequest, t reflect.Type) reflect.Value {
	if a.UnmarshalConversationToken == nil {
		return reflect.Zero(t)
	}
	f := reflect.ValueOf(a.UnmarshalConversationToken)
	str := req.GetConversationToken()
	arg0 := reflect.ValueOf([]byte(str))
	vs := f.Call([]reflect.Value{arg0})
	if str != "" && !vs[1].IsNil() {
		log.Warningf(ctx, "faild to get conversationToken: %v", vs[1])
	}
	return vs[0]
}

func (a *AppHandler) setConversationToken(ctx context.Context, req *api.AppRequest, res *api.AppResponse, v reflect.Value) {
	if a.MarshalConversationToken == nil {
		return
	}
	f := reflect.ValueOf(a.MarshalConversationToken)
	vs := f.Call([]reflect.Value{v})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to set conversationToken: %v", vs[1])
		return
	}
	res.ConversationToken = string(vs[0].Bytes())
}

func (a *AppHandler) getUserStorage(ctx context.Context, req *api.AppRequest, t reflect.Type) reflect.Value {
	if a.UnmarshalUserStorage == nil {
		return reflect.Zero(t)
	}
	f := reflect.ValueOf(a.UnmarshalUserStorage)
	str := req.GetUserStorage()
	arg0 := reflect.ValueOf([]byte(str))
	vs := f.Call([]reflect.Value{arg0})
	if str != "" && !vs[1].IsNil() {
		log.Warningf(ctx, "faild to get conversationToken: %v", vs[1])
	}
	return vs[0]
}

func (a *AppHandler) setUserStorage(ctx context.Context, req *api.AppRequest, res *api.AppResponse, v reflect.Value) {
	if a.MarshalUserStorage == nil {
		return
	}
	f := reflect.ValueOf(a.MarshalUserStorage)
	vs := f.Call([]reflect.Value{v})
	if !vs[1].IsNil() {
		log.Warningf(ctx, "faild to set conversationToken: %v", vs[1])
		return
	}
	res.UserStorage = string(vs[0].Bytes())
}

func (a *AppHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	ctx = context.WithValue(ctx, KEY_URL_PATH, r.URL.Path)
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
	f := reflect.ValueOf(a.HandleApp)
	tokenType := f.Type().In(3)
	storageType := f.Type().In(4)
	token := a.getConversationToken(ctx, &req, tokenType)
	storage := a.getUserStorage(ctx, &req, storageType)
	arg0 := reflect.ValueOf(ctx)
	arg1 := reflect.ValueOf(&req)
	arg2 := reflect.ValueOf(&res)
	f.Call([]reflect.Value{arg0, arg1, arg2, token, storage})
	a.setConversationToken(ctx, &req, &res, token)
	a.setUserStorage(ctx, &req, &res, storage)
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf(ctx, "response body: %s", js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UrlPath(ctx context.Context) string {
	value := ctx.Value(KEY_URL_PATH)
	str, _ := value.(string)
	return str
}
