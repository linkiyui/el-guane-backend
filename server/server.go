package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	api_server "gitlab.com/hospital-web/api_server"
)

func Start() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Println("terminating CNXXXN server | signal: ", s)
		os.Exit(0)
	}()

	api_server.NewApiServer().Start()
}
