package database

import (
	"errors"
	"time"
)

type History struct {
	Response 	string 		`json:"response"`
	Duration 	float32		`json:"duration"`
	CreatedAt	int64 		`json:"created_at"`
}

type Request struct {
	ID          int    				`json:"id"`
	Url       	string 				`json:"url"`
	Interval 	time.Duration 		`json:"interval"`
	History		[]History 			`json:"-"`
}


type AllRequests []Request
var Requests = AllRequests{}

func GetById(id int) (Request, error) {
	for i, singleEvent := range Requests {
		if singleEvent.ID == id {
			return Requests[i], nil
		}
	}
	return Request{}, errors.New("Cannot find the resource")
}
