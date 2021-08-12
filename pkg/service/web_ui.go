package service

import (
	"context"
	"fmt"
	"github.com/JungleMC/web-ui/internal/config"
	"github.com/JungleMC/web-ui/internal/pages"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Instance *WebUiService

type WebUiService struct {
}

func Start() {
	config.Get = &config.Config{}
	err := env.Parse(config.Get)
	if err != nil {
		panic(err)
	}

	Instance = &WebUiService{}
	Instance.Bootstrap()
}

func (s *WebUiService) Bootstrap() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	r.HandleFunc("/", pages.Home)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%v:%v", config.Get.WebHost, config.Get.WebPort),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	_ = srv.Shutdown(ctx)
	os.Exit(0)
}
