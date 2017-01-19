package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type User struct {
	Name string
}

func postOne(req *restful.Request, resp *restful.Response) {
	newUser := new(User)
	err := req.ReadEntity(newUser)
	if err != nil {
		//resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("new user: '%v'", newUser)
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").To(postOne).
		Param(ws.BodyParameter("User", "A User").DataType("main.User")))

	restful.Add(ws)

	http.ListenAndServe(":8080", nil)
}

/*
import ("io"
	"net/http"
	//"ConfigurationRepository"
)

func cfhService(res http.ResponseWriter, req *http.Request){
	io.WriteString(res,"Care Free Home Central Service running")
}

func main(){
	http.HandleFunc("/", cfhService)
	http.ListenAndServe(":8080",nil)
	http.HandleFunc("/savenewuser", saveUser)
}

func saveUser(res http.ResponseWriter, req *http.Request){
	//ConfigurationRepository.SaveUser(req)
	io.WriteString(res,"this is res")
}*/