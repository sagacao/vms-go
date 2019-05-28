package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
	"vms-go/goweb/dbs"
	"vms-go/goweb/routers"
	"vms-go/goweb/settings"

	"github.com/sagacao/goworld/engine/binutil"
	"github.com/sagacao/goworld/engine/gwlog"

	"github.com/gin-gonic/gin"
)

func prepare() {
	// logger.SetRollingDaily(settings.SvrConfig.Env.ACCESS_LOG_PATH, settings.SvrConfig.Env.ACCESS_LOG_NAME)
	logfile := path.Join(settings.SvrConfig.Env.ACCESS_LOG_PATH, settings.SvrConfig.Env.ACCESS_LOG_NAME)
	if settings.SvrConfig.Env.DEBUG {
		gin.SetMode(gin.DebugMode)
		// logger.SetConsole(true)
		// logger.SetLevel(logger.DEBUG)
		binutil.SetupGWLog("vms", "debug", logfile, true)
	} else {
		gin.SetMode(gin.ReleaseMode)
		// logger.SetLevel(logger.INFO)
		binutil.SetupGWLog("vms", "info", logfile, false)
	}

	err := dbs.NewDBService()
	if err != nil {
		gwlog.Fatal("mysql init ", err)
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
			gwlog.Fatal("listen: ", err)
		}
	}()
	// logger.Info("Http Serving at: ", addr)
	gwlog.Infof("Http Serving at: %s", addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	gwlog.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		gwlog.Fatal("Server Shutdown:", err)
	}
	gwlog.Infof("Server exiting")
}
