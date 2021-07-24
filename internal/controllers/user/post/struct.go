package post

// Input Data
type RequestPostUser struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// Output Data
type ResponsePostUser struct {
	UUID string `json:"uuid"`
}
