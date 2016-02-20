package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/k0kubun/pp"
	"github.com/kardianos/osext"
	"github.com/mattn/go-pipeline"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	pp.Fprintf(w, "<h1>Hello, Electron-Go!</h1>")
}

func main() {
	flag.Set("bind", ":8080")
	goji.Get("/", hello)
	go goji.Serve()

	err := launchElectron()

	if err == nil {
		terminate(0)
	} else {
		log.Fatal(err)
		terminate(1)
	}

	// Termination.
	defer func() {
		terminate(0)
	}()

	return
}

func terminate(code int) {
	graceful.ShutdownNow()
	os.Exit(code)
}

func launchElectron() error {
	// Get current path
    var folderPath string
    var err error

    cond := "run" // go build の時は何か適当なのに変える（ダサ...）
    var elec, elecarg string
    if cond == "run" {
        // launch directory
        folderPath, err = os.Getwd()
        if err != nil {
            return err
        }
        elec = "electron"
        elecarg = "./Electron"
    } else {
                // launch binary
        folderPath, err = osext.ExecutableFolder()
        if err != nil {
            return err
        }
        elec = folderPath + "/Electron"
        elecarg = ""
    }

	out, err := pipeline.Output(
		[]string{elec, elecarg},
	)
	if err != nil {
		pp.Println(err)
		return err
	}
	pp.Println(string(out))
	return nil
}
