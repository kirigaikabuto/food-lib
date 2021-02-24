package food_lib

type Food struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Price       int64    `json:"price"`
}
