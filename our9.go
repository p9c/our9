package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/our9/rts"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
	// comhttpus                     = "com-http.us"

	comhttpus   = "com-http.us"
	bitnodesnet = "bitnodes.net"
)

var r = mux.NewRouter()

func main() {
	////////////////
	// COM-HTTP.US
	////////////////
	r.Headers("X-Requested-With", "XMLHttpRequest")

	r.Host(comhttpus).Path("/").HandlerFunc(rts.ComHttpUs).Name("comhttpus")
	r.Host(comhttpus).Path("/coins/").HandlerFunc(rts.Coins).Methods("GET")
	r.Host(comhttpus).Path("/bitnodes/").HandlerFunc(rts.BitNodes).Methods("GET")
	r.Host(comhttpus).Path("/bitnodes/{coin}/{nodeid}").HandlerFunc(rts.BitNode).Methods("GET")

	r.Host("{coin}." + comhttpus).Path("/").HandlerFunc(rts.Coin).Methods("GET")
	r.Host("{coin}." + comhttpus).Path("/favicon.ico").HandlerFunc(rts.Ico).Name("ico")

	r.Host("{coin}." + comhttpus).Path("/network").HandlerFunc(rts.Network).Methods("GET")
	r.Host("{coin}." + comhttpus).Path("/community").HandlerFunc(rts.Community).Methods("GET")
	r.Host("{coin}." + comhttpus).Path("/art").HandlerFunc(rts.Art).Methods("GET")

	r.Host("{coin}." + comhttpus).Path("/explorer").HandlerFunc(rts.Explorer).Methods("GET")
	r.Host("{coin}." + comhttpus).Path("/explorer/block/{block}").HandlerFunc(rts.Block).Name("block")
	r.Host("{coin}." + comhttpus).Path("/explorer/hash/{hash}").HandlerFunc(rts.Hash).Name("hash")
	r.Host("{coin}." + comhttpus).Path("/explorer/tx/{tx}").HandlerFunc(rts.Tx).Name("tx")
	// r.Host("{coin}." + comhttpus).Path("/explorer/address/{id}").HandlerFunc(rts.Address).Name("address")
	r.Host("{coin}." + comhttpus).Path("/explorer/search").HandlerFunc(rts.Search).Name("search")

	// // r.HandleFunc("/", rts.AmpFrontHandler) // GET
	// // r.HandleFunc("/coins/{slug}/f/cmc", rts.CMCHandler).Name("cmc")

	r.Host("{coin}." + comhttpus).Path("/api/b").HandlerFunc(rts.ApiLastBlock).Name("b")
	r.Host("{coin}." + comhttpus).Path("/api/bta/{id}").HandlerFunc(rts.ApiBlockTxAddr).Name("bta")

	r.Host("api." + comhttpus).Path("/bitnodes").HandlerFunc(rts.ApiBitNodes).Name("bitnodes")

	r.Host("api." + comhttpus).Path("/{coin}/i").HandlerFunc(rts.ApiInfo).Name("info")
	r.Host("api." + comhttpus).Path("/{coin}/p").HandlerFunc(rts.ApiPeer).Name("peer")
	r.Host("api." + comhttpus).Path("/{coin}/m").HandlerFunc(rts.ApiMiningInfo).Name("mining")
	r.Host("api." + comhttpus).Path("/{coin}/r").HandlerFunc(rts.ApiRawMemPool).Name("rawmempool")

	r.Host("api." + comhttpus).Path("/{coin}/n").HandlerFunc(rts.ApiNodes).Name("nodes")
	r.Host("api." + comhttpus).Path("/{coin}/n/{nodeid}").HandlerFunc(rts.ApiBitNode).Name("nodes")

	r.Host("api." + comhttpus).Path("/{coin}/{type}/{id}").HandlerFunc(rts.ApiData).Name("coin")

	r.Host("i." + comhttpus).Path("/{coin}/{size}").HandlerFunc(rts.Img).Name("img")
	r.Host("f." + comhttpus).Path("/{coin}/{type}/{file}").HandlerFunc(rts.Frame).Name("frames")
	r.Host("l." + comhttpus).Path("/{type}/{file}").HandlerFunc(rts.Libs).Name("libs")

	////////////////
	// COM-HTTP.US
	////////////////
	r.Host(bitnodesnet).Path("/").HandlerFunc(rts.BitNodesNet).Name("bitnodesnet")
	r.Host(bitnodesnet).Path("/amp").HandlerFunc(rts.BitNodesNetAmp).Name("bitnodesnetamp")
	r.Host(bitnodesnet).Path("/amp/co").HandlerFunc(rts.BitNodesNetAmpCo).Name("bitnodesnetampco")
	// r.Host(bitnodesnet).Path("/amp/pay").HandlerFunc(rts.BitNodesNetApiPay).Name("bitnodesnetapipay")
	r.Host(bitnodesnet).Path("/api/pay").HandlerFunc(rts.BitNodesNetApiPayCall).Name("bitnodesnetapipaycall")

	r.Host(bitnodesnet).Path("/s/{file}").HandlerFunc(rts.BitNodesNetStatic).Name("static")

	// // r.Schemes("https")

	// go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(r)))

	srv := &http.Server{
		Handler: handlers.CORS()(handlers.CompressHandler(r)),
		// Handler: cacheHandler(handlers.CORS()(handlers.CompressHandler(r))),
		// Handler: handlers.CompressHandler(r),
		Addr: ":443",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}

// func cacheHandler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var age time.Duration
// 		ext := filepath.Ext(r.URL.String())

// 		// Timings are based on github.com/h5bp/server-configs-nginx

// 		switch ext {
// 		case ".rss", ".atom":
// 			age = time.Hour / time.Second
// 		case ".css", ".js":
// 			age = (time.Hour * 24 * 365) / time.Second
// 		case ".jpg", ".jpeg", ".gif", ".png", ".ico", ".cur", ".gz", ".svg", ".svgz", ".mp4", ".ogg", ".ogv", ".webm", ".htc":
// 			age = (time.Hour * 24 * 30) / time.Second
// 		default:
// 			age = 1
// 		}

// 		if age > 0 {
// 			w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", age))
// 			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
// 		}

// 		h.ServeHTTP(w, r)
// 	})
// }
