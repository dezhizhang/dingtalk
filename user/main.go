package main

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

func main() {
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	go func() {
		log.Fatalf("user server running in %s\n", srv.Addr)

		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("shutting down user web server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("user server shutdown,cause by", err)
		return
	}

	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}

	fmt.Println("user server stop success...")

}
