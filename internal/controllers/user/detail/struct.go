package detail

type RequestDetailUser struct {
	UUID string `json:"uuid"`
}

type ResponseDetailUser struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
