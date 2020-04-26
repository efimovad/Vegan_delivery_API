package dishusecase

import (
	"github.com/efimovad/Vegan_delivery_API/internal/app/dish"
	"github.com/efimovad/Vegan_delivery_API/internal/models"
)

type Usecase struct {
}

func NewDishUsecase() dish.IUsecase {
	return &Usecase{
	}
}

func (u * Usecase) GetDishes(cafeID int64) ([]models.Dish, error) {
	dish1 := models.Dish{
		ID:          1,
		Name:        "Грибная 31 см",
		Ingredients: "шампиньоны, веган сыр, мукаб оливковое масло",
		Cost:        635,
		Image:       "https://sun9-43.userapi.com/impg/c206628/v206628939/88a46/BNCik2ojUZQ.jpg?size=520x0&quality=90&sign=42c1ae1e039bf6633d104ed93d01544b",
	}

	dish2 := models.Dish{
		ID:          2,
		Name:        "Пепперони 31 см",
		Ingredients: "пшеничные сосиски, веган сыр, мукаб оливковое масло",
		Cost:        700,
		Image:       "https://sun9-25.userapi.com/impg/c855136/v855136939/1f7c92/0UD6cIRn1tQ.jpg?size=520x0&quality=90&sign=8906f656ebac01da6c24826a3ef21720",
	}

	dish3 := models.Dish{
		ID:          3,
		Name:        "Чилинтано 35 см",
		Ingredients: "томатный соус, растительное мясо «Green wise», халапеньо,чили, соевая моцарелла",
		Cost:        740,
		Image:       "https://sun9-57.userapi.com/impf/c857120/v857120149/67/OfCAIWZhrbk.jpg?size=520x0&quality=90&sign=94f41ad7d13a50ed89db688172b15ef6",
	}

	dish4 := models.Dish{
		ID:          4,
		Name:        "Маргарита 35 см",
		Ingredients: "томатный соус, базилик, томаты, соевая моцарелла",
		Cost:        740,
		Image:       "https://sun9-15.userapi.com/impg/c858216/v858216939/18dde8/BvVwBVnKAw0.jpg?size=520x0&quality=90&sign=a7cfbfd2027251b24c72c22f97fb84b1",
	}

	return []models.Dish{dish1, dish2, dish3, dish4}, nil
}

func (u * Usecase) GetDish(ID int64) (models.Dish, error) {
	return models.Dish{
		ID:          4,
		Name:        "Маргарита 35 см",
		Ingredients: "томатный соус, базилик, томаты, соевая моцарелла",
		Cost:        740,
		Image:       "https://sun9-15.userapi.com/impg/c858216/v858216939/18dde8/BvVwBVnKAw0.jpg?size=520x0&quality=90&sign=a7cfbfd2027251b24c72c22f97fb84b1",
	}, nil
}

func (u * Usecase) AddDish(dish models.Dish) error {
	panic("implement me")
}