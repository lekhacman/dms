package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal/handler"
	"github.com/lekhacman/dms/internal/store"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func Index(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprint(ctx, "Document Center Service")
}

func main() {

	db, err := store.NewDmsStore(store.DbSpec{
		"postgres",
		os.Getenv("dbpass"),
		"dms",
		"localhost",
		"5432",
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT * FROM documents")
	if err != nil {
		log.Fatal(err)
	}

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/model", handler.Model)
	router.POST("/", handler.Create)

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}
