package bar

type Bar struct {
	ID int `json:"id"`
}

type BarResponse struct {
	Data []Bar `json:"data"`
}
