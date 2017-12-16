package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	"github.com/valyala/fasthttp"
)

// Open a text file with the excuses
var file, _ = ioutil.ReadFile("./excuses.txt")

// Split strings
var strs = strings.Split(string(file), "\n")

func main() {
	h := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/excuse":
			getRandomExcuse(ctx)
		default:
			ctx.Error("Page Not Found", fasthttp.StatusNotFound)
		}

		log.Print("Connection established at :", ctx.ConnTime(), ", IP: ", ctx.RemoteIP())
	}

	s := &fasthttp.Server{
		Handler: h,
		Name:    "Excuse",
	}
	log.Fatal(s.ListenAndServe(":8080"))

}

// Returns a random excuse
// from a text file
func getRandomExcuse(ctx *fasthttp.RequestCtx) {

	random := strs[rand.Intn(len(strs))]
	data := struct {
		Msg string `json:"msg"`
	}{
		random,
	}
	j, _ := json.Marshal(data)
	ctx.SetContentType("application/json")
	ctx.SetBody(j)
}
