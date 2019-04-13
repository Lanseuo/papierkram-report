package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	zipFilepath := flag.String("filepath", "", "ZIP-File exported from papierkram.de")
	port := flag.Int("port", 8181, "Port for webserver")

	flag.Parse()

	err := unzip(*zipFilepath)
	if err != nil {
		log.Fatalln("Unable to unzip file", err)
	}

	parseData()
	startServer(*port)
}

func startServer(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/api", apiHandler).Methods("GET")
	router.HandleFunc("/api/balance", balanceHandler).Methods("GET")
	router.HandleFunc("/api/balance/development", balanceDevelopmentHandler).Methods("GET")
	router.PathPrefix("/").HandlerFunc(staticFilesHandler)

	handler := cors.Default().Handler(router)

	fmt.Printf("Server is running on port %d: http://localhost:%d\n", port, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatalf("Cannot start http server on port %d\n", port)
	}

}

func staticFilesHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if url == "/" {
		url = "/index.html"
	}

	if strings.HasSuffix(url, ".html") {
		w.Header().Set("Content-Type", "text/html")
	} else if strings.HasSuffix(url, ".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(url, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}

	bytes, err := Asset("static" + url)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not find " + url))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(bytes)
}

func unzip(zipFile string) error {
	destination := "/tmp/papierkram-report"

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(destination, f.Name)

		// Check for ZipSlip
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
