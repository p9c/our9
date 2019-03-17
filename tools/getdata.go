package tools

import (
	"io"

	// "compress/gzip"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
)

type proxyPage struct {
	Time time.Time `json:"time"`
	Data []byte    `json:"data"`
}

var PxP = make(map[string]proxyPage)

// func getData(r *http.Request) (*TelemetryEvents, error) {
// 	var data TelemetryEvents
// 	var decoder *json.Decoder
// 	switch r.Header.Get("Content-Encoding") {
// 	case "gzip":
// 		gz, err := gzip.NewReader(r.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer gz.Close()
// 		decoder = json.NewDecoder(gz)
// 	default:
// 		decoder = html. json.NewDecoder(r.Body)
// 	}
// 	err := decoder.Decode(&data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func GetData(w http.ResponseWriter, r *http.Request, page, url string) {
	gdb, err := OpenDB()
	if err != nil {
	}
	gPage := proxyPage{}
	gdb.Read("cache", "pages", &gPage)
	pTime := PxP[page].Time
	timeNow := time.Now()
	pTp := pTime.Add(time.Duration(1 * time.Minute))
	if timeNow.After(pTp) {
		gData, err := http.Get(url)
		if err != nil {
			fmt.Println("Get Proxy Dat Fail", gData)
		}
		defer gData.Body.Close()
		pD, err := ioutil.ReadAll(gData.Body)

		PxP[page] = proxyPage{
			Time: timeNow,
			Data: pD,
		}

		gdb.Write("cache", "pages", PxP)
	}
	// w.Write([]byte(PxP[page].Data))
	// return template.HTML(fmt.Sprint(string(PxP[page].Data)))

	io.WriteString(w, string(PxP[page].Data))
}

func GetDataDirect(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)
	return mapBody, err
}

func OpenDB() (jdb *scribble.Driver, err error) {
	jdb, err = scribble.New("/web/comhttp/jdb", nil)
	if err != nil {
		fmt.Println("ErrorDB", err)
	}
	return
}
