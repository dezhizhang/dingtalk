package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Fatalf("%s server running in %s\n", srvName, srv.Addr)

		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Printf("shutting down %s web server...", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s server shutdown,cause by", err)
		return
	}

	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}

	fmt.Printf("%s server stop success...", srvName)
}
