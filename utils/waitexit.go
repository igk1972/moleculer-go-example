package utils

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type process interface {
	Stop() (err error)
}

func WaitExit(serve process) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		sig := <-sigs
		Log.Info("on system signal: ", sig)
		Log.Debug(sig)
		done <- true
	}()
	Log.Info("wait for SIGINT/SIGTERM ...")
	<-done
	serve.Stop()
	time.Sleep(time.Second * 1)
	os.Exit(0)
}
