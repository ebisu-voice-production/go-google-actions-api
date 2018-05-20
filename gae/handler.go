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

type GaeHandler struct {
	MakeConversationToken func() interface{}
	MakeUserStorage       func() interface{}
	HandleApp             func(ctx context.Context, req *api.AppRequest, res *api.AppResponse, conversationToken interface{}, userStorage interface{})
}

func (g *GaeHandler) getConversationToken(ctx context.Context, req *api.AppRequest) interface{} {
	if g.MakeConversationToken == nil {
		return nil
	}
	value := g.MakeConversationToken()
	err := json.Unmarshal([]byte(req.GetConversationToken()), &value)
	if err != nil {
		log.Warningf(ctx, "faild to get conversationToken: %v", err)
	}
	return value
}

func (g *GaeHandler) setConversationToken(ctx context.Context, value interface{}, req *api.AppRequest, res *api.AppResponse) {
	if value == nil {
		return
	}
	b, err := json.Marshal(value)
	if err != nil {
		log.Warningf(ctx, "faild to set conversationToken: %v", err)
		return
	}
	res.ConversationToken = string(b)
}

func (g *GaeHandler) getUserStorage(ctx context.Context, req *api.AppRequest) interface{} {
	if g.MakeUserStorage == nil {
		return nil
	}
	value := g.MakeUserStorage()
	err := json.Unmarshal([]byte(req.GetUserStorage()), &value)
	if err != nil {
		log.Warningf(ctx, "faild to parse userStorage: %v", err)
	}
	return value
}

func (g *GaeHandler) setUserStorage(ctx context.Context, value interface{}, req *api.AppRequest, res *api.AppResponse) {
	if value == nil {
		return
	}
	b, err := json.Marshal(value)
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
	conversationToken := g.getConversationToken(ctx, &req)
	userStorage := g.getUserStorage(ctx, &req)
	if conversationToken == nil && userStorage == nil {
		g.HandleApp(ctx, &req, &res, nil, nil)
	} else if conversationToken == nil {
		g.HandleApp(ctx, &req, &res, nil, &userStorage)
	} else if userStorage == nil {
		g.HandleApp(ctx, &req, &res, &conversationToken, nil)
	} else {
		g.HandleApp(ctx, &req, &res, &conversationToken, &userStorage)
	}
	g.setConversationToken(ctx, conversationToken, &req, &res)
	g.setUserStorage(ctx, userStorage, &req, &res)
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf(ctx, "response body: %s", js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
