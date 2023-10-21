package person

type PersonDTO struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Age     *int32  `json:"age"`
	Address *string `json:"address"`
	Work    *string `json:"work"`
}
