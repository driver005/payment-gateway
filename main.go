package main

import (
	"context"
	"embed"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	// "github.com/gofiber/fiber/v2/middleware/limiter"
)

// opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
// sh := middleware.Redoc(opts, nil)

// public.Handler("GET", "/docs", sh)
// public.Handler("GET", "/swagger.json", http.FileServer(http.Dir("./")))

//go:embed migrations/*.sql
var migrations embed.FS

var ctx = context.Background()

// func CreateInvoice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	handler := btcpay.NewHandler()
// 	handler.SetClient("https://testnet.demo.btcpayserver.org", "d0b9689840047ed7b0737d9e981ab5e83c1b9d35")
// 	handler.SetStoreID()

// 	invoice := handler.CreateInvoice()

// 	respond.NewResponse(w).Ok(invoice)
// }

// func Webhook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
// 	// 	resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	// 	return
// 	// }
// 	body, _ := ioutil.ReadAll(r.Body)
// 	fmt.Println(string(body))
// 	_ = ioutil.WriteFile("./data.txt", body, 0666)
// }

// func TrackDerivationScheme(c *gin.Context) {
// 	// if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
// 	// 	resp.NewResponse(w).InternalServerError(helper.WithStack(err))
// 	// 	return
// 	// }
// 	// fmt.Println(r.URL.Query().Get("name"))
// 	// fmt.Println(reflect.TypeOf(r.URL.Query().Get("name")))
// }

// func ParseField(fieldStruct reflect.StructField) {
// 	indirectFieldType := fieldStruct.Type
// 	for indirectFieldType.Kind() == reflect.Ptr {
// 		indirectFieldType = indirectFieldType.Elem()
// 	}
// 	fieldValue := reflect.New(indirectFieldType)
// 	fmt.Println(reflect.Indirect(fieldValue).Kind())
// 	switch reflect.Indirect(fieldValue).Kind() {
// 	case reflect.Struct:
// 		fmt.Println(fieldValue.Interface())
// 	}
// }

// func createQuery(q interface{}) {
// 	if reflect.ValueOf(q).Kind() == reflect.Struct {
// 		t := reflect.TypeOf(q).Name()
// 		query := fmt.Sprintf("insert into %s values(", t)
// 		v := reflect.ValueOf(q)
// 		// rs := reflect.ValueOf(&q).Elem()
// 		for i := 0; i < v.NumField(); i++ {
// 			rf := v.Field(i)
// 			fieldName := v.Type().Field(i).Name
// 			switch v.Field(i).Kind() {
// 			case reflect.Int:
// 				if i == 0 {
// 					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
// 				} else {
// 					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
// 				}
// 			case reflect.String:
// 				if i == 0 {
// 					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
// 				} else {
// 					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
// 				}
// 			case reflect.Struct:
// 				// fmt.Println(valueField.Kind())
// 				if fieldName != "Filter" {
// 					createQuery(v.Field(i).Elem().Interface())
// 				}

// 				modelType := reflect.ValueOf(rf).Type()
// 				for modelType.Kind() == reflect.Slice || modelType.Kind() == reflect.Array || modelType.Kind() == reflect.Ptr {
// 					modelType = modelType.Elem()
// 				}
// 				for i := 0; i < modelType.NumField(); i++ {
// 					fieldStruct := modelType.Field(i)
// 					fmt.Println(fieldStruct.Name)
// 					ParseField(fieldStruct)
// 				}
// 			case reflect.Pointer:
// 				rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr()))
// 				// ri.Set(rf)
// 				// rf.Set(ri)

// 				fmt.Println(rf)
// 			default:
// 				fmt.Println("Unsupported type")
// 				return
// 			}
// 		}
// 		query = fmt.Sprintf("%s)", query)
// 		fmt.Println(query)
// 		return

// 	}
// 	fmt.Println("unsupported type")

// }

// func DeepFields(iface interface{}) []reflect.Value {
// 	fields := make([]reflect.Value, 0)
// 	ifv := reflect.ValueOf(iface)
// 	ift := reflect.TypeOf(iface)

// 	for i := 0; i < ift.NumField(); i++ {
// 		v := ifv.Field(i)

// 		switch v.Kind() {
// 		case reflect.Struct:
// 			fields = append(fields, DeepFields(v.Interface())...)
// 		default:
// 			fields = append(fields, v)
// 		}
// 	}

// 	return fields
// }

// func collectFieldNames(iface interface{}, m map[string]struct{}) {
// 	t := reflect.TypeOf(iface)
// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}
// 	if t.Kind() != reflect.Struct {
// 		return
// 	}
// 	for i := 0; i < t.NumField(); i++ {
// 		sf := t.Field(i)
// 		m[sf.Name] = struct{}{}
// 		if sf.Anonymous {
// 			collectFieldNames(sf.Type, m)
// 		}
// 	}
// }

// func BuildQuery(i []*structs.Field) string {
// 	var query string
// 	for _, f := range i {
// 		name := strings.Replace(f.Tag("json"), ",omitempty", "", -1)
// 		if f.Kind() == reflect.Struct {
// 			if f.Name() == "Lt" {
// 				if !f.IsZero() {
// 					query += fmt.Sprintf(`< "%+v"`, f.Value())
// 				}
// 			} else if f.Name() == "Gt" {
// 				if !f.IsZero() {
// 					query += fmt.Sprintf(`> "%+v"`, f.Value())
// 				}
// 			} else if f.Name() == "Lte" {
// 				if !f.IsZero() {
// 					query += fmt.Sprintf(`<= "%+v"`, f.Value())
// 				}
// 			} else if f.Name() == "Gte" {
// 				if !f.IsZero() {
// 					query += fmt.Sprintf(`>= "%+v"`, f.Value())
// 				}
// 			} else {
// 				if !f.IsZero() {
// 					if f.IsEmbedded() {
// 						query = BuildQuery(f.Fields())
// 					} else {
// 						if len(query) > 0 {
// 							query += " AND "
// 						}
// 						query += fmt.Sprintf("%+v %+v", name, BuildQuery(f.Fields()))
// 					}

// 				}
// 			}
// 		} else if f.Kind() == reflect.String {
// 			if len(query) > 0 {
// 				query += " AND "
// 			}
// 			if !f.IsZero() {
// 				query += fmt.Sprintf(`%+v = "%+v"`, name, f.Value())
// 			}
// 		} else if f.Kind() == reflect.Int {
// 			if len(query) > 0 {
// 				query += " AND "
// 			}
// 			if !f.IsZero() {
// 				query += fmt.Sprintf(`%+v = %+v`, name, f.Value())
// 			}
// 		} else if f.Kind() == reflect.Bool {
// 			if len(query) > 0 {
// 				query += " AND "
// 			}
// 			if !f.IsZero() {
// 				query += fmt.Sprintf(`%+v = %+v`, name, f.Value())
// 			}
// 		}
// 	}
// 	return query
// }

func main() {
	// r := driver.New(ctx)
	// m := models.Product{}

	// err := r.Repository().GetWhere(&m, "id = '6c1b336c-156a-4dcd-a124-0a143a07117c'")
	// fmt.Println(err)
	// fmt.Println(m)

	// repository.BuildFindQuery(repository.FindOption{
	// 	Select: map[string]interface{}{
	// 		"firstName": true,
	// 		"lastName":  false,
	// 	},
	// })

	// f := types.FilterableBatchJobProps{}
	// f.CreatedAt.Lt = time.Now().UTC().Round(time.Second)
	// f.UpdatedAt.Gt = time.Now().UTC().Round(time.Second)
	// f.Offset = 2

	// id, _ := uuid.NewUUID()
	// m := models.BatchJob{}
	// m.Id = id

	// r := driver.New(ctx)
	// q := r.ClientManager().Query(ctx, f, m)

	// stmt := q.Where("id = ?", "test").First(&models.BatchJob{})

	// fmt.Println(stmt.Statement.SQL.String())
	// fmt.Println(stmt.Statement.Vars...)

	// r := driver.New(ctx)

	// r.Setup()

	// p := &models.Product{
	// 	Title: "Shirt",
	// }

	// stmt := r.Manager(context.Background()).Session(&database.Session{DryRun: true}).Create(p)
	// if stmt.Error != nil {
	// 	fmt.Println(stmt.Error)
	// }
	// fmt.Println(stmt.Statement.SQL.String())
	// fmt.Println(stmt.Statement.Vars)

	// public := fiber.New(fiber.Config{
	// 	ServerHeader:  "Fiber",
	// 	AppName:       "Test App v1.0.1",
	// 	WriteTimeout:  15 * time.Second,
	// 	ReadTimeout:   15 * time.Second,
	// 	StrictRouting: true,
	// })

	// public.Use(cors.New())
	// // public.Use(csrf.New())
	// public.Use(logger.New())
	// // public.Use(limiter.New())

	// r.RegisterRoutes(public)

	// public.Listen("192.168.1.2:80")

	// var m models.ShippingProfile

	// if err := r.Manager(ctx).Where("type = ?", "gift_card").First(&m).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// e := r.Manager(context.Background()).Create(&models.Product{
	// 	IsGiftcard: true,
	// 	Title:      "test",
	// 	ProfileId:  uuid.NullUUID{UUID: m.Id, Valid: true},
	// })

	// if pgError := e.Error.(*pq.Error); errors.Is(e.Error, pgError) {
	// 	// switch pgError.Code {
	// 	// case "23505":

	// 	// }
	// 	fmt.Println(pgError.Hint)
	// 	fmt.Println(pgError.Detail)
	// }

	// conn, err := pgx.Connect(context.Background(), "postgres://jzyozqtc:Hdh78SBKXkvgs6-Z5fpVQAw_la4Iln__@batyr.db.elephantsql.com/jzyozqtc")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// var name string
	// var weight int64
	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(name, weight)

	// r.Handler().Routes(public)

	// n := handler.NewNbxplorer(r, "http://localhost:32838")

	// public.GET("/checkout", CreateInvoice)
	// public.POST("/webhook", Webhook)

	// public.GET("/test", TrackDerivationScheme)

	// r.RegisterRoutes(public)
}

// Server Configuration

// srv := &http.Server{
// 	Handler:      public,
// 	Addr:         "192.168.1.2:80",
// 	WriteTimeout: 15 * time.Second,
// 	ReadTimeout:  15 * time.Second,
// }
// fmt.Println("Starting server")
// defer srv.Close()
// log.Fatal(srv.ListenAndServe())

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
