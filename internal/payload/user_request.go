package payload

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
