package entities

// Input Date
// DS:Date Structure
type User struct {
	UUID string `db:"uuid"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

func New(uuid string, name string, age int64) *User {
	return &User{
		UUID: uuid,
		Name: name,
		Age:  age,
	}
}
