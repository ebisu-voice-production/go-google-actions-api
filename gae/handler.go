package gae

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/ebisu-voice-production/go-google-actions-api/api"
)

type AppContext interface {
	GetConversationToken() interface{}
	GetUserStorage() interface{}
	HandleApp(ctx context.Context, req *api.AppRequest, res *api.AppResponse)
}

type GaeHandler struct {
	MakeAppContext func() AppContext
}

func (g *GaeHandler) getConversationToken(ctx context.Context, appCtx AppContext, req *api.AppRequest) {
	x := appCtx.GetConversationToken()
	if x == nil {
		return
	}
	err := json.Unmarshal([]byte(req.GetConversationToken()), &x)
	if err != nil {
		log.Warningf(ctx, "faild to get conversationToken: %v", err)
	}
}

func (g *GaeHandler) setConversationToken(ctx context.Context, appCtx AppContext, req *api.AppRequest, res *api.AppResponse) {
	x := appCtx.GetConversationToken()
	if x == nil {
		return
	}
	b, err := json.Marshal(x)
	if err != nil {
		log.Warningf(ctx, "faild to set conversationToken: %v", err)
		return
	}
	res.ConversationToken = string(b)
}

func (g *GaeHandler) getUserStorage(ctx context.Context, appCtx AppContext, req *api.AppRequest) {
	x := appCtx.GetUserStorage()
	if x == nil {
		return
	}
	err := json.Unmarshal([]byte(req.GetUserStorage()), &x)
	if err != nil {
		log.Warningf(ctx, "faild to parse userStorage: %v", err)
	}
}

func (g *GaeHandler) setUserStorage(ctx context.Context, appCtx AppContext, req *api.AppRequest, res *api.AppResponse) {
	x := appCtx.GetUserStorage()
	if x == nil {
		return
	}
	b, err := json.Marshal(x)
	if err != nil {
		log.Warningf(ctx, "faild to set userStorage: %v", err)
		return
	}
	newUserStorage := string(b)
	if req.GetUserStorage() != newUserStorage {
		res.UserStorage = newUserStorage
	}
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
	appCtx := g.MakeAppContext()
	g.getConversationToken(ctx, appCtx, &req)
	g.getUserStorage(ctx, appCtx, &req)
	appCtx.HandleApp(ctx, &req, &res)
	g.setConversationToken(ctx, appCtx, &req, &res)
	g.setUserStorage(ctx, appCtx, &req, &res)
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf(ctx, "response body: %s", js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
