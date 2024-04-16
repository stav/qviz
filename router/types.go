package router

// Contains the common types used in the router package

type Quiz struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Date string   `json:"created_at"`
	Ques Question `json:"question"`
	Qcnt int      `json:"count"`
	Msg  string
}

func (q *Quiz) Error() string {
	return q.Msg
}

type Question struct {
	ID   int      `json:"id"`
	Num  int      `json:"number"`
	Text string   `json:"text"`
	Ans  []Answer `json:"answer"`
}

type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	True bool   `json:"is_correct"`
}
