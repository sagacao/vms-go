package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vms-go/goweb/dbs"
	"vms-go/goweb/logger"
	"vms-go/goweb/routers"
	"vms-go/goweb/settings"

	"github.com/gin-gonic/gin"
)

func prepare() {
	logger.SetRollingDaily(settings.SvrConfig.Env.ACCESS_LOG_PATH, settings.SvrConfig.Env.ACCESS_LOG_NAME)
	if settings.SvrConfig.Env.DEBUG {
		gin.SetMode(gin.DebugMode)
		logger.SetConsole(true)
		logger.SetLevel(logger.DEBUG)
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger.SetLevel(logger.INFO)
	}

	err := dbs.NewDBService()
	if err != nil {
		logger.Fatal("mysql init ", err)
	}
}

func main() {
	prepare()
	router := routers.InitRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		//logger.Debug("origin", origin)
		//w.Header().Set("Access-Control-Allow-Origin", "*") --http cors -cooki
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		router.ServeHTTP(w, r)
	})

	addr := settings.SvrConfig.Env.SERVER_ADDR //":8080"

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: ", err)
		}
	}()
	logger.Info("Http Serving at: ", addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
}
