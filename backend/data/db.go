package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//DBPublicationList структура для базы данных
type DBPublicationList struct {
	dbpub *gorm.DB
}

//DBProfileList структура для базы данных
type DBProfileList struct {
	db *gorm.DB
}

//DBCommentList структура для базы данных
type DBCommentList struct {
	dbcom *gorm.DB
}

//NewDBPublicationList база данных
func NewDBPublicationList(connection string) (*DBPublicationList, error) {
	dbpub, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	dbpub.AutoMigrate(&Publication{})
	return &DBPublicationList{dbpub: dbpub}, dbpub.Error
}

//ClosePublication закрытие
func (cl *DBPublicationList) ClosePublication() {
	cl.dbpub.Close()
}

//GetPublications запрос GET
func (cl *DBPublicationList) GetPublications() []Publication {
	var publications []Publication
	cl.dbpub.Find(&publications)
	return publications
}

//AddPublication запрос POST
func (cl *DBPublicationList) AddPublication(publication Publication) int {
	cl.dbpub.Create(&publication)
	return publication.IDpub
}

//EditPublication запрос PUT
func (cl *DBPublicationList) EditPublication(idpub int, publication Publication) error {
	var publications []Publication
	cl.dbpub.Where("idpub = ?", idpub).Find(&publications)
	if len(publications) == 0 {
		return fmt.Errorf("can't find record #%d", idpub)
	}

	publication.IDpub = publications[0].IDpub
	cl.dbpub.Save(&publication)
	return cl.dbpub.Error
}

//RemovePublication запрос DELETE
func (cl *DBPublicationList) RemovePublication(idpub int) error {
	var publications []Publication
	cl.dbpub.Where("idpub = ?", idpub).Find(&publications)
	if len(publications) == 0 {
		return fmt.Errorf("can't find record #%d", idpub)
	}
	cl.dbpub.Delete(&publications[0])
	return cl.dbpub.Error
}

//NewDBProfileList бд
func NewDBProfileList(connection string) (*DBProfileList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Profile{})
	return &DBProfileList{db: db}, db.Error
}

//CloseProfile закрытие
func (cl *DBProfileList) CloseProfile() {
	cl.db.Close()
}

//GetProfiles запрос GET
func (cl *DBProfileList) GetProfiles() []Profile {
	var profiles []Profile
	cl.db.Find(&profiles)
	return profiles
}

//AddProfile запрос POST
func (cl *DBProfileList) AddProfile(profile Profile) int {
	cl.db.Create(&profile)
	return profile.ID
}

//EditProfile запрос PUT
func (cl *DBProfileList) EditProfile(id int, profile Profile) error {
	var profiles []Profile
	cl.db.Where("id = ?", id).Find(&profiles)
	if len(profiles) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	profile.ID = profiles[0].ID
	cl.db.Save(&profile)
	return cl.db.Error
}

//RemoveProfile запрос DELETE
func (cl *DBProfileList) RemoveProfile(id int) error {
	var profiles []Profile
	cl.db.Where("id = ?", id).Find(&profiles)
	if len(profiles) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&profiles[0])
	return cl.db.Error
}

//NewDBCommentList бд
func NewDBCommentList(connection string) (*DBCommentList, error) {
	dbcom, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	dbcom.AutoMigrate(&Comment{})
	return &DBCommentList{dbcom: dbcom}, dbcom.Error
}

//CloseComment закрытие
func (cl *DBCommentList) CloseComment() {
	cl.dbcom.Close()
}

//GetComments запрос GET
func (cl *DBCommentList) GetComments() []Comment {
	var сomments []Comment
	cl.dbcom.Find(&сomments)
	return сomments
}

//AddComment запрос POST
func (cl *DBCommentList) AddComment(comment Comment) int {
	cl.dbcom.Create(&comment)
	return comment.IDcom
}

//EditComment запрос PUT
func (cl *DBCommentList) EditComment(idcom int, comment Comment) error {
	var comments []Comment
	cl.dbcom.Where("idcom = ?", idcom).Find(&comments)
	if len(comments) == 0 {
		return fmt.Errorf("can't find record #%d", idcom)
	}

	comment.IDcom = comments[0].IDcom
	cl.dbcom.Save(&comment)
	return cl.dbcom.Error
}

//RemoveComment запрос DELETE
func (cl *DBCommentList) RemoveComment(idcom int) error {
	var comments []Comment
	cl.dbcom.Where("idcom = ?", idcom).Find(&comments)
	if len(comments) == 0 {
		return fmt.Errorf("can't find record #%d", idcom)
	}
	cl.dbcom.Delete(&comments[0])
	return cl.dbcom.Error
}
