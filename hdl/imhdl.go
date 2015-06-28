package hdl

import (
	//	"github.com/cihub/seelog"
	"net/http"
	"github.com/lankecheng/polyge/server/serv/im"
	"github.com/lankecheng/polyge/server/serv"
)

func WebsocketConnect(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	oauthToken := req.FormValue("token")
	ok, err := serv.CheckToken(oauthToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	uid, err := serv.GetUidFromToken(oauthToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = im.EstablishConnect(uid, w, req)
}