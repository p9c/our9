package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parallelcointeam/our9/cfg"
	"github.com/parallelcointeam/our9/tools"

	"github.com/gorilla/mux"
)

func ApiData(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/" + mux.Vars(r)["type"] + "/" + mux.Vars(r)["id"]
	tools.GetData(w, r, mux.Vars(r)["coin"]+"Data"+mux.Vars(r)["type"]+mux.Vars(r)["id"], url)
}
func ApiLast(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/last"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"Last", url)
}

func ApiLastBlock(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/b"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"LastBlock", url)
}

func ApiBlockTxAddr(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/b/" + mux.Vars(r)["id"]
	tools.GetData(w, r, mux.Vars(r)["coin"]+"BTA"+mux.Vars(r)["id"], url)
}

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/info"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"Info", url)
}
func ApiPeer(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/peer"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"Peer", url)
}
func ApiMiningInfo(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/gmi"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"MiningInfo", url)
}
func ApiRawMemPool(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/rmp"
	tools.GetData(w, r, mux.Vars(r)["coin"]+"RawMemPool", url)
}

func ApiNodes(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/n/" + mux.Vars(r)["coin"]
	tools.GetData(w, r, mux.Vars(r)["coin"]+"Nodes", url)
}
func ApiBitNode(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/n/" + mux.Vars(r)["coin"] + "/" + mux.Vars(r)["nodeid"]
	tools.GetData(w, r, mux.Vars(r)["coin"]+"BitNode"+mux.Vars(r)["nodeid"], url)
}
func ApiBitNodes(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Jorm + "/n"
	tools.GetData(w, r, "BitNodesStat", url)
}

func DoSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	search := r.FormValue("chsrc")
	fmt.Println("searchsearchsearchsearchsearchsearchsearchsearch", search)

	tps := []string{"block", "hash", "tx", "addr"}
	var tpt = "noresults"
	for _, tp := range tps {
		url := "http://" + cfg.Jorm + "/e/" + mux.Vars(r)["coin"] + "/" + tp + "/" + search
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", url)
		resp, err := tools.GetDataDirect(url)
		var search map[string]interface{}
		json.Unmarshal(resp, &search)
		if err != nil {
			fmt.Println("Read error", err)
		}
		if search["d"] != nil {
			tpt = tp

		}

	}

	http.Redirect(w, r, fmt.Sprintf("/explorer/"+tpt+"/"+search), http.StatusPermanentRedirect)
}
