package database

import "errors"

type Request struct {
	ID          int    	`json:"id"`
	Url       	string 	`json:"url"`
	Interval 	int 	`json:"interval"`
}


type AllRequests []Request

var Events = AllRequests{}

func GetById(id int) (Request, error) {
	for i, singleEvent := range Events {
		if singleEvent.ID == id {
			return Events[i], nil
		}
	}
	return Request{}, errors.New("Cannot find the resource")
}
