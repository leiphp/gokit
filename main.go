package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"gokit/initialize"
	. "gokit/services"
	"golang.org/x/time/rate"
	kitlog "github.com/go-kit/kit/log"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main () {
	//服务参数获取
	name := flag.String("name","","服务名称")
	port := flag.Int("p",0,"服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定端口")
	}
	initialize.SetServiceNameAndPort(*name, *port) //设置服务名和端口
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger,"gokit","1.0")
		logger = kitlog.With(logger,"time",kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger,"caller",kitlog.DefaultCaller)
	}

	user := UserService{}	//用户服务
	limit := rate.NewLimiter(1, 5)
	endp := RateLimit(limit)(UserServiceLogMiddleware(logger)(CheckTokenMiddleware()(GenUserEndpoint(user))))

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(MyErrorEncoder),
	}

	serverHandler := httptransport.NewServer(endp,DecodeUserRequest,EncodeUserRequest,options...)

	//增加handler用于获取用户token
	accessService := &AccessService{}
	accessServiceEndpoint := AccessEndpoint(accessService)
	accessHandler := httptransport.NewServer(accessServiceEndpoint,DecodeAccessRequest,EncodeAccessRequest,options...)


	router := mymux.NewRouter()
	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	router.Methods("POST").Path("/access-token").Handler(accessHandler)
	router.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	router.Methods("GET").Path("/health").HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		write.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go func() {
		//注册consul服务
		initialize.RegisterServer()
		err := http.ListenAndServe(":"+strconv.Itoa(*port),router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s",<-sig_c)
	}()

	//如果没有异常错误，errChan将永久阻塞
	getErr := <- errChan
	initialize.UnregisterServer()
	log.Println(getErr)
}