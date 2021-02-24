package food_lib

type ListCommand struct {
}

func (cmd *ListCommand) Exec(service FoodService) (interface{}, error) {
	return service.List(cmd)
}

type CreateCommand struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Price       int64    `json:"price"`
}

func (cmd *CreateCommand) Exec(service FoodService) (interface{}, error) {
	return service.Create(cmd)
}
