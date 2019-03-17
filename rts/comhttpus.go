package rts

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/parallelcointeam/our9/cfg"
	"github.com/parallelcointeam/our9/tools"
)

func ComHttpUs(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/"
	tools.GetData(w, r, "home", url)

}
func Coins(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/"
	tools.GetData(w, r, "coins", url)
}
func BitNodes(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/bitnodes/"
	tools.GetData(w, r, "bitnodes", url)
}
func BitNode(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/bitnodes/" + mux.Vars(r)["coin"] + "/" + mux.Vars(r)["nodeid"]
	tools.GetData(w, r, "bitnodes"+mux.Vars(r)["coin"]+mux.Vars(r)["nodeid"], url)
}
func Coin(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"]
	tools.GetData(w, r, "coin"+mux.Vars(r)["coin"], url)
}
func Ico(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/favicon.ico"
	tools.GetData(w, r, "coin"+mux.Vars(r)["coin"]+"ico", url)
}
func Network(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/network"
	tools.GetData(w, r, "network"+mux.Vars(r)["coin"], url)
}
func Community(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/community"
	tools.GetData(w, r, "community"+mux.Vars(r)["coin"], url)
}
func Art(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/art"
	tools.GetData(w, r, "art"+mux.Vars(r)["coin"], url)
}
func Explorer(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/explorer"
	tools.GetData(w, r, "explorer"+mux.Vars(r)["coin"], url)
}

func Block(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/explorer/block/" + mux.Vars(r)["block"]
	tools.GetData(w, r, "block"+mux.Vars(r)["coin"]+mux.Vars(r)["block"], url)
}

func Hash(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/explorer/hash/" + mux.Vars(r)["hash"]
	tools.GetData(w, r, "hash"+mux.Vars(r)["coin"]+mux.Vars(r)["hash"], url)
}

func Tx(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/explorer/tx/" + mux.Vars(r)["tx"]
	tools.GetData(w, r, "tx"+mux.Vars(r)["coin"]+mux.Vars(r)["tx"], url)
}

func Search(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/explorer/search/" + mux.Vars(r)["search"]
	tools.GetData(w, r, "search"+mux.Vars(r)["coin"]+mux.Vars(r)["search"], url)
}
func Img(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/img/" + mux.Vars(r)["size"]
	tools.GetData(w, r, "img"+mux.Vars(r)["coin"]+mux.Vars(r)["size"], url)
}
func Frame(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/coins/" + mux.Vars(r)["coin"] + "/frame/" + mux.Vars(r)["type"] + "/" + mux.Vars(r)["file"]
	tools.GetData(w, r, "frame"+mux.Vars(r)["coin"]+mux.Vars(r)["type"]+mux.Vars(r)["file"], url)
}
func Libs(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.Enso + "/libs/" + mux.Vars(r)["type"] + "/" + mux.Vars(r)["file"]
	tools.GetData(w, r, "libs"+mux.Vars(r)["coin"]+mux.Vars(r)["type"]+mux.Vars(r)["file"], url)
}
