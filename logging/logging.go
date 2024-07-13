package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("log with microseconds")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file")

	myLog := log.New(os.Stdout, "submodule:", log.LstdFlags)
	myLog.Println("from myLog")

	myLog.SetPrefix("custom:")
	myLog.Println("from myLog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}
