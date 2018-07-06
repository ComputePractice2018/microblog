package data

import "fmt"

//Publication структура для хранения публикации
type Publication struct {
	Namepub    string `json:"namepub"`
	Time       string `json:"time"`
	Publcation string `json:"publication"`
}

//publications хранимый список публикаций
var publications []Publication

//GetPublications возвращает список публикаций
func GetPublications() []Publication {
	return publications
}

//AddPublication добавляет публикацию в конец списка и возвращает id
func AddPublication(publication Publication) int {
	id := len(publications)
	publications = append(publications, publication)
	return id
}

//EditPublication изменяет публикацию с id на publication
func EditPublication(publication Publication, id int) error {
	if id < 0 || id >= len(publications) {
		return fmt.Errorf("icorrect ID")
	}
	publications[id] = publication
	return nil
}

//RemovePublication удаляет публикацию по id
func RemovePublication(id int) error {
	if id < 0 || id >= len(publications) {
		return fmt.Errorf("icorrect ID")
	}
	publications = append(publications[:id], publications[id+1:]...)
	return nil
}
