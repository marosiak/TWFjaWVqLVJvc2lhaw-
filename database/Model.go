package database

type Request struct {
	ID          int    	`json:"id"`
	Url       	string 	`json:"url"`
	Interval 	int 	`json:"interval"`
}


type AllRequests []Request

var Events = AllRequests{}

func GetLength() int {
	return len(AllRequests{})
}