package food_lib

type FoodStore interface {
	List() ([]Food, error)
	Create(food *Food) (*Food, error)
}
