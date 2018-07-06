package data

//Publication структура для хранения публикации
type Publication struct {
	Namepub    string `json:"namepub"`
	Time       string `json:"time"`
	Publcation string `json:"publication"`
}

//PublicationList хранимый список публикаций
var PublicationList = []Publication{Publication{
	Namepub:    "Название",
	Time:       "Время",
	Publcation: "Публикация"}}
