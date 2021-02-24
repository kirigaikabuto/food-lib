package food_lib

import (
	"database/sql"
	"encoding/json"
	"log"
	_ "github.com/lib/pq"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS foods (
		id serial,
		name text,
		description text,
		price int,
		image text,
		images text,
		PRIMARY KEY(id)
	);`,
}

type postgreFoodStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (FoodStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreFoodStore{db: db}, err
}

func (foodStore *postgreFoodStore) List() ([]Food, error) {
	var foods []Food
	data, err := foodStore.db.Query("select id, name, description, price, image, images from foods")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		food := Food{}
		images := ""
		err = data.Scan(&food.Id, &food.Name, &food.Description, &food.Price, &food.Image, &images)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(images), &food.Images)
		foods = append(foods, food)
	}
	return foods, nil
}

func (foodStore *postgreFoodStore) Create(food *Food) (*Food, error) {
	images, err := json.Marshal(food.Images)
	if err != nil {
		return nil, err
	}
	err = foodStore.db.QueryRow("insert into foods (name, description, price, image, images) values ($1,$2,$3,$4,$5) RETURNING id", food.Name, food.Description, food.Price, food.Image, images).Scan(&food.Id)
	if err != nil {
		return nil, err
	}
	return food, nil
}
