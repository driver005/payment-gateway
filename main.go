package main

import (
	"context"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/driver005/gateway/btcpay"
	"github.com/driver005/gateway/driver"
	"github.com/driver005/gateway/helper"
	"github.com/driver005/gateway/respond"
	"github.com/julienschmidt/httprouter"
)

// opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
// sh := middleware.Redoc(opts, nil)

// public.Handler("GET", "/docs", sh)
// public.Handler("GET", "/swagger.json", http.FileServer(http.Dir("./")))

//go:embed migrations/*.sql
var migrations embed.FS

var ctx = context.Background()

func CreateInvoice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	handler := btcpay.NewHandler()
	handler.SetClient("https://testnet.demo.btcpayserver.org", "d0b9689840047ed7b0737d9e981ab5e83c1b9d35")
	handler.SetStoreID()

	invoice := handler.CreateInvoice()

	respond.NewResponse(w).Ok(invoice)
}

func Webhook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
	// 	resp.NewResponse(w).InternalServerError(helper.WithStack(err))
	// 	return
	// }
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	_ = ioutil.WriteFile("./data.txt", body, 0666)
}

func TrackDerivationScheme(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
	// 	resp.NewResponse(w).InternalServerError(helper.WithStack(err))
	// 	return
	// }
	fmt.Println(r.URL.Query().Get("name"))
	fmt.Println(reflect.TypeOf(r.URL.Query().Get("name")))
}

func main() {
	// mux := http.NewServeMux()
	r := driver.New(ctx)

	public := helper.NewRouter()
	r.RegisterRoutes(public)
	// r.Handler().Routes(public)

	// n := handler.NewNbxplorer(r, "http://localhost:32838")

	// public.GET("/checkout", CreateInvoice)
	// public.POST("/webhook", Webhook)

	// public.GET("/test", TrackDerivationScheme)

	// r.RegisterRoutes(public)

	// Server Configuration

	srv := &http.Server{
		Handler:      public,
		Addr:         "192.168.1.2:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server")
	defer srv.Close()
	log.Fatal(srv.ListenAndServe())
}

// func newLogger() *zap.Logger {
// 	return logger.NewZapLogger(&config.Logger{})
// }

// func newEthRPCClient() *ethrpc.Client {
// 	var err error
// 	rpcEthClient, err := ethereum.NewRPCClient(&config.Ethereum{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return rpcEthClient
// }
// ethe, _ := ethereum.NewEthereum(
// 	newEthRPCClient(),
// 	&config.Ethereum{},
// 	newLogger(),
// 	"eth",
// )

// tags := []eth.QuantityTag{
// 	eth.QuantityTagLatest,
// 	eth.QuantityTagPending,
// 	// eth.QuantityTagEarliest,
// }

// type args struct {
// 	addr string
// }
// type want struct {
// 	balance uint64
// }
// tests := []struct {
// 	name string
// 	args args
// 	want want
// }{
// 	{
// 		name: "happy path",
// 		args: args{addr: "0xEA247646137F74C38e04f3012db483d77F3dEc59"},
// 		want: want{50000000000000000},
// 	},
// 	{
// 		name: "happy path",
// 		args: args{addr: "0x6B91314E559D3FB5b40f9F51582631a7b5C610ef"},
// 		want: want{50000000000000000},
// 	},
// 	{
// 		name: "happy path",
// 		args: args{addr: "0x0aC5c95EB979C41CFa2C6BdF8e5515F966fEc103"},
// 		want: want{50000000000000000},
// 	},
// }
// for _, tt := range tests {

// 	for _, tag := range tags {
// 		balance, err := ethe.GetBalance(tt.args.addr, tag)

// 		if err == nil {
// 			fmt.Printf("quantityTag: %s, balance: %d", tag, balance.Uint64())
// 		}
// 	}

// }

// r := driver.New(ctx)
