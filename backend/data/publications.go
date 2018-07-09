package data

import "fmt"

//Publication структура для хранения публикации
type Publication struct {
	Namepub    string `json:"namepub"`
	Time       string `json:"time"`
	Publcation string `json:"publication"`
}

//Profile структура для храниения профиля пользователя
type Profile struct {
	Nikname string `json:"nikname"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Github  string `json:"github"`
}

//Comment структура для хранения комментария
type Comment struct {
	Time    string `json:"time"`
	Comment string `json:"comment"`
}

//publications хранимый список публикаций
var publications []Publication

//profiles хранимый список профилей
var profiles []Profile

//comments хранимый список комментариев
var comments []Comment

//AddProfile добавляет профиль в конец списка и возвращает id
func AddProfile(profile Profile) int {
	id := len(profiles)
	profiles = append(profiles, profile)
	return id
}

//EditProfile изменяет публикацию с id на profile
func EditProfile(profile Profile, id int) error {
	if id < 0 || id >= len(profiles) {
		return fmt.Errorf("icorrect ID")
	}
	profiles[id] = profile
	return nil
}

//GetPublications возвращает список публикаций
func GetPublications() []Publication {
	return publications
}

//AddPublication добавляет публикацию в конец списка и возвращает idpub
func AddPublication(id int, publication Publication) int {
	idpub := len(publications)
	publications = append(publications, publication)
	return idpub
}

//EditPublication изменяет публикацию с id на publication
func EditPublication(id int, publication Publication, idpub int) error {
	if id < 0 || id >= len(profiles) {
		return fmt.Errorf("icorrect ID")
	}

	if idpub < 0 || idpub >= len(publications) {
		return fmt.Errorf("icorrect ID publication")
	}
	publications[idpub] = publication
	return nil
}

//RemovePublication удаляет публикацию по idpub
func RemovePublication(id int, idpub int) error {
	if id < 0 || id >= len(profiles) {
		return fmt.Errorf("icorrect ID")
	}

	if idpub < 0 || idpub >= len(publications) {
		return fmt.Errorf("icorrect ID publication")
	}
	publications = append(publications[:idpub], publications[idpub+1:]...)
	return nil
}

//GetComments возвращает список комментариев
func GetComments(id int, idpub int) []Comment {
	return comments
}

//AddComment добавляет публикацию в конец списка и возвращает id
func AddComment(id int, idpub int, comment Comment) int {
	idcom := len(comments)
	comments = append(comments, comment)
	return idcom
}

//EditComment изменяет публикацию с id на comment
func EditComment(id int, idpub int, comment Comment, idcom int) error {
	if idpub < 0 || idpub >= len(publications) {
		return fmt.Errorf("icorrect ID publication")
	}

	if idcom < 0 || idcom >= len(comments) {
		return fmt.Errorf("icorrect ID comment")
	}
	comments[idcom] = comment
	return nil
}

//RemoveComment удаляет публикацию по id
func RemoveComment(id int, idpub int, idcom int) error {
	if idpub < 0 || idpub >= len(publications) {
		return fmt.Errorf("icorrect ID publication")
	}

	if idcom < 0 || idcom >= len(comments) {
		return fmt.Errorf("icorrect ID comment")
	}
	comments = append(comments[:idcom], comments[idcom+1:]...)
	return nil
}
