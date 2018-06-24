package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/wuriyanto48/go-rabbitmq-example/producer/internal/pub"
	"github.com/wuriyanto48/go-rabbitmq-example/producer/internal/utils"
)

//HTTPHandler struct
type HTTPHandler struct {
	q string
	publisher  pub.Publisher
}

//NewHTTPHandler function
func NewHTTPHandler(q string, publisher pub.Publisher) *HTTPHandler {
	return &HTTPHandler{q: q, publisher: publisher}
}

//PublishMessage handler function
func (h *HTTPHandler) PublishMessages() http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){

		if req.Method != "POST" {
			log.Println("Invalid Method")
			utils.JSONResponse(res, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}

		var message pub.Message

		//get message from request
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&message)

		if err != nil {
			log.Printf("Error parsing message : %v", err.Error())
			utils.JSONResponse(res, "Invalid Method", http.StatusBadRequest)
			return
		}

		//publish message to rabbitmq
		b, err := message.JSON()

		if err != nil {
			log.Printf("Error %s", err.Error())
			utils.JSONResponse(res, "Error occured", http.StatusInternalServerError)
			return
		}

		err = h.publisher.Publish(h.q, b)

		if err != nil {
			log.Printf("Error %s", err.Error())
			utils.JSONResponse(res, "Error occured", http.StatusInternalServerError)
			return
		}

		utils.JSONResponse(res, "Message sent", http.StatusOK)
	})
}