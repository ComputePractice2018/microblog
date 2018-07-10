package data

import "fmt"

//Publication структура для хранения публикации
type Publication struct {
	Namepub    string `json:"namepub"`
	Time       string `json:"time"`
	Publcation string `json:"publication"`
}

//PublicationList структура для списка публикаций
type PublicationList struct {
	publications []Publication
}

//EditablePub интерфейс для работы со списком публикаций
type EditablePub interface {
	GetPublications() []Publication
	AddPublication(publication Publication) int
	EditPublication(idpub int, publication Publication) error
	RemovePublication(idpub int) error
}

//NewPublicationList конструктор списка публикаций
func NewPublicationList() *PublicationList {
	return &PublicationList{}
}

//Profile структура для храниения профиля
type Profile struct {
	Nikname string `json:"nikname"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Github  string `json:"github"`
}

//ProfileList структура для списка профилей
type ProfileList struct {
	profiles []Profile
}

//EditableProfile интерфейс для работы со списком профиля
type EditableProfile interface {
	GetProfiles() []Profile
	AddProfile(profile Profile) int
	EditProfile(id int, profile Profile) error
	RemoveProfile(id int) error
}

//NewProfileList конструктор списка профиля
func NewProfileList() *ProfileList {
	return &ProfileList{}
}

//Comment структура для хранения комментария
type Comment struct {
	Time    string `json:"time"`
	Comment string `json:"comment"`
}

//CommentList структура для списка комментариев
type CommentList struct {
	comments []Comment
}

//EditableComment интерфейс для работы со списком комментариев
type EditableComment interface {
	GetComments() []Comment
	AddComment(comment Comment) int
	EditComment(idcom int, comment Comment) error
	RemoveComment(idcom int) error
}

//NewCommentList конструктор списка комментариев
func NewCommentList() *CommentList {
	return &CommentList{}
}

//GetProfiles возвращает список профилей
func (cl *ProfileList) GetProfiles() []Profile {
	return cl.profiles
}

//AddProfile добавляет профиль в конец списка и возвращает id
func (cl *ProfileList) AddProfile(profile Profile) int {
	id := len(cl.profiles)
	cl.profiles = append(cl.profiles, profile)
	return id
}

//EditProfile изменяет профиль с id на profile
func (cl *ProfileList) EditProfile(id int, profile Profile) error {
	if id < 0 || id >= len(cl.profiles) {
		return fmt.Errorf("icorrect ID")
	}
	cl.profiles[id] = profile
	return nil
}

//RemoveProfile удаляет профиль по idpub
func (cl *ProfileList) RemoveProfile(id int) error {
	if id < 0 || id >= len(cl.profiles) {
		return fmt.Errorf("icorrect ID")
	}
	cl.profiles = append(cl.profiles[:id], cl.profiles[id+1:]...)
	return nil
}

//GetPublications возвращает список публикаций
func (cl *PublicationList) GetPublications() []Publication {
	return cl.publications
}

//AddPublication добавляет публикацию в конец списка и возвращает idpub
func (cl *PublicationList) AddPublication(publication Publication) int {
	idpub := len(cl.publications)
	cl.publications = append(cl.publications, publication)
	return idpub
}

//EditPublication изменяет публикацию с id на publication
func (cl *PublicationList) EditPublication(idpub int, publication Publication) error {

	if idpub < 0 || idpub >= len(cl.publications) {
		return fmt.Errorf("icorrect ID publication")
	}
	cl.publications[idpub] = publication
	return nil
}

//RemovePublication удаляет публикацию по idpub
func (cl *PublicationList) RemovePublication(idpub int) error {

	if idpub < 0 || idpub >= len(cl.publications) {
		return fmt.Errorf("icorrect ID publication")
	}
	cl.publications = append(cl.publications[:idpub], cl.publications[idpub+1:]...)
	return nil
}

//GetComments возвращает список комментариев
func (cl *CommentList) GetComments() []Comment {
	return cl.comments
}

//AddComment добавляет комментарий в конец списка и возвращает id
func (cl *CommentList) AddComment(comment Comment) int {
	idcom := len(cl.comments)
	cl.comments = append(cl.comments, comment)
	return idcom
}

//EditComment изменяет комментарий с id на comment
func (cl *CommentList) EditComment(idcom int, comment Comment) error {

	if idcom < 0 || idcom >= len(cl.comments) {
		return fmt.Errorf("icorrect ID comment")
	}
	cl.comments[idcom] = comment
	return nil
}

//RemoveComment удаляет комментарий по id
func (cl *CommentList) RemoveComment(idcom int) error {

	if idcom < 0 || idcom >= len(cl.comments) {
		return fmt.Errorf("icorrect ID comment")
	}
	cl.comments = append(cl.comments[:idcom], cl.comments[idcom+1:]...)
	return nil
}
