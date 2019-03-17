package rts

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our9/cfg"
	"github.com/parallelcointeam/our9/tools"
)

func BitNodesNet(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.BitNodes + "/"

	tools.GetData(w, r, "bitnodesnet", url)

}
func BitNodesNetAmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://bitnodes.net")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "https://bitnodes.net")

	url := "http://" + cfg.BitNodes + "/amp"
	tools.GetData(w, r, "amp", url)

}
func BitNodesNetAmpCo(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.BitNodes + "/amp/co"
	tools.GetData(w, r, "ampco", url)

}

// func BitNodesNetApiPay(w http.ResponseWriter, r *http.Request) {
// 	url := "http://" + cfg.BitNodes + "/api/pay"
// 	tools.GetData(w, r, "apipay", url)

// }
func BitNodesNetApiPayCall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://bitnodes.net")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "https://bitnodes.net")

	url := "http://" + cfg.BitNodes + "/api/pay"
	tools.GetData(w, r, "apipay", url)

}

// func Coins(w http.ResponseWriter, r *http.Request) {
// 	url := "http://" + cfg.Enso + "/coins/"
// 	tools.GetData(w, r, "coins", url)
// }
// func BitNodes(w http.ResponseWriter, r *http.Request) {
// 	url := "http://" + cfg.Enso + "/bitnodes/"
// 	tools.GetData(w, r, "bitnodes", url)
// }
// func BitNode(w http.ResponseWriter, r *http.Request) {
// 	url := "http://" + cfg.Enso + "/bitnodes/" + mux.Vars(r)["coin"] + "/" + mux.Vars(r)["nodeid"]
// 	tools.GetData(w, r, "bitnodes"+mux.Vars(r)["coin"]+mux.Vars(r)["nodeid"], url)
// }
func BitNodesNetStatic(w http.ResponseWriter, r *http.Request) {
	url := "http://" + cfg.BitNodes + "/s/" + mux.Vars(r)["file"]

	tools.GetData(w, r, "static"+mux.Vars(r)["file"], url)

}
