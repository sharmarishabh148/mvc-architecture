package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sharmari/mvc/services"
	"github.com/sharmari/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	//userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprint("user id must be an integer !"),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.MarshalIndent(apiErr, "", "  ")
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	log.Printf("about to process user_id %v", userId)
	log.Printf("return the User to the client")
	//v := reflect.ValueOf(user)
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
