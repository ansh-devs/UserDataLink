package entity

type User struct {
	ID      int     `json:"id"`
	FName   string  `json:"fName"`
	Height  float64 `json:"Height"`
	City    string  `json:"City"`
	Phone   string  `json:"Phone"`
	Married bool    `json:"Married"`
}
