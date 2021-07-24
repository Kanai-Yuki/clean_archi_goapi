package entities

// Input Date
// DS:Date Structure
type User struct {
	UUID string
	Name string
	Age  int64
}

func New(uuid string, name string, age int64) *User {
	return &User{
		UUID: uuid,
		Name: name,
		Age:  age,
	}
}
