package data

//Publications структура для хранения публикации
type Publications struct {
	Namepub    string `json:"namepub"`
	Time       string `json:"time"`
	Publcation string `json:"publication"`
}

//PublicationList хранимый список публикаций
var PublicationList = []Publications{Publications{
	Namepub:    "Название",
	Time:       "Время",
	Publcation: "Публикация"}}
