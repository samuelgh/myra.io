package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"encoding/json"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {

			//w.WriteJson(map[string]string{"Body": "Hello World!"})
			n := make(map[string]string)

			data := req.URL.Query()
			for key, val := range data {
				log.Print("key:" + key)
				flatString := strings.Join(val,",")
				n[key] = flatString
			}
			log.Print(n)
			jsonString,_ := json.Marshal(n)
			print(jsonString);
			w.WriteJson(jsonString)
			//w.WriteJson(readItems2);
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}