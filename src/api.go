package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
)

func InitAPI() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/device/:name", func(w rest.ResponseWriter, req *rest.Request) {
			name := req.PathParam("name")
			n := make(map[string]string)
			data := req.URL.Query()
			for key, val := range data {
				log.Println("key:" + key)
				flatString := strings.Join(val,",")
				n[key] = flatString
			}
			item := Item{}
			err := item.FillStruct(n)
			if err != nil {
				fmt.Println(err)
			}
			result := Store(item,name)
			fmt.Println(item)
			log.Print(n)
			jsonString,_ := json.Marshal(&n)
			print(jsonString);
			w.WriteJson(result)
		}),
		rest.Get("/device/:name/all", func(w rest.ResponseWriter, req *rest.Request) {
			id := req.PathParam("name")
			meta := ReadMetaByName(id)
			data := ReadByName(id)
			result := ResultData{meta, data}
			w.WriteJson(result)
		}),
		rest.Post("/device/:name", func(w rest.ResponseWriter, req *rest.Request) {
			newItem := ItemMeta{}
			req.DecodeJsonPayload(&newItem)
			StoreMetaData(newItem)
			CreateTableWithName(req.PathParam("name"))
			w.WriteHeader(200)
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