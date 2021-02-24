package food_lib

type FoodService interface {
	List(cmd *ListCommand) ([]Food, error)
	Create(cmd *CreateCommand) (*Food, error)
}

type foodService struct {
	foodStore FoodStore
}

func NewProductService(foodStore FoodStore) FoodService {
	return &foodService{foodStore: foodStore}
}

func (svc *foodService) List(cmd *ListCommand) ([]Food, error) {
	return svc.foodStore.List()
}

func (svc *foodService) Create(cmd *CreateCommand) (*Food, error) {
	return svc.foodStore.Create(&Food{
		Name:        cmd.Name,
		Description: cmd.Description,
		Image:       cmd.Image,
		Images:      cmd.Images,
		Price:       cmd.Price,
	})
}
