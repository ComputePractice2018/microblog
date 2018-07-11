package data

import "testing"

var testProfiles = []Profile{
	{
		ID:      1,
		Nikname: "Alx",
		Name:    "Alexandr",
		Surname: "Nikishin",
		Email:   "San4a-97@mail.ru",
		Github:  "Sanalx",
	},
	{
		ID:      2,
		Nikname: "Ivan",
		Name:    "Ivan",
		Surname: "Ivanov",
		Email:   "Ivanov@mail.ru",
		Github:  "Iva",
	},
}

var testPublications = []Publication{
	{
		IDpub:      1,
		Namepub:    "Backend",
		Time:       "21:00",
		Publcation: "Run backend test...",
	},
	{
		IDpub:      2,
		Namepub:    "Frontend",
		Time:       "00:00",
		Publcation: "Run frontend test...",
	},
}

var testComments = []Comment{
	{
		IDcom:   1,
		Time:    "12:00",
		Comment: "Alexandr",
	},
	{
		IDcom:   2,
		Time:    "01:10",
		Comment: "Ivan",
	},
}

func TestAddPublication(t *testing.T) {
	cl := NewPublicationList()

	cl.AddPublication(testPublications[0])

	if cl.GetPublications()[0] != testPublications[0] {
		t.Errorf("AddPublication is not working")
	}
}

func TestEditPublication(t *testing.T) {
	cl := NewPublicationList()

	cl.AddPublication(testPublications[0])

	err := cl.EditPublication(1, testPublications[1])

	if cl.GetPublications()[0] != testPublications[1] {
		t.Errorf("EditPublication is not working")
	}
	if err != nil {
		t.Errorf("Unexpented EditPublication error: %+v", err)
	}

	err = cl.EditPublication(-1, testPublications[1])

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestDeletePublication(t *testing.T) {
	cl := NewPublicationList()
	cl.AddPublication(testPublications[0])
	cl.AddPublication(testPublications[1])

	err := cl.RemovePublication(1)

	if cl.GetPublications()[0] != testPublications[1] {
		t.Errorf("RemovePublication is not working")
	}
	if err != nil {
		t.Errorf("Unexpented RemovePublication error")
	}

	err = cl.RemovePublication(-1)

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestAddProfiles(t *testing.T) {
	cl := NewProfileList()

	cl.AddProfile(testProfiles[0])

	if cl.GetProfiles()[0] != testProfiles[0] {
		t.Errorf("AddProfile is not working")
	}
}

func TestEditProfile(t *testing.T) {
	cl := NewProfileList()

	cl.AddProfile(testProfiles[0])

	err := cl.EditProfile(1, testProfiles[1])

	if cl.GetProfiles()[0] != testProfiles[1] {
		t.Errorf("EditProfiles is not working")
	}
	if err != nil {
		t.Errorf("Unexpented EditProfiles error: %+v", err)
	}

	err = cl.EditProfile(-1, testProfiles[1])

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestDeleteProfiles(t *testing.T) {
	cl := NewProfileList()
	cl.AddProfile(testProfiles[0])
	cl.AddProfile(testProfiles[1])

	err := cl.RemoveProfile(1)

	if cl.GetProfiles()[0] != testProfiles[1] {
		t.Errorf("RemoveProfile is not working")
	}
	if err != nil {
		t.Errorf("Unexpented RemoveProfile error")
	}

	err = cl.RemoveProfile(-1)

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestAddComments(t *testing.T) {
	cl := NewCommentList()

	cl.AddComment(testComments[0])

	if cl.GetComments()[0] != testComments[0] {
		t.Errorf("AddComment is not working")
	}
}

func TestEditComment(t *testing.T) {
	cl := NewCommentList()

	cl.AddComment(testComments[0])

	err := cl.EditComment(1, testComments[1])

	if cl.GetComments()[0] != testComments[1] {
		t.Errorf("EditComments is not working")
	}
	if err != nil {
		t.Errorf("Unexpented EditComments error: %+v", err)
	}

	err = cl.EditComment(-1, testComments[1])

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestDeleteComments(t *testing.T) {
	cl := NewCommentList()
	cl.AddComment(testComments[0])
	cl.AddComment(testComments[1])

	err := cl.RemoveComment(1)

	if cl.GetComments()[0] != testComments[1] {
		t.Errorf("RemoveComment is not working")
	}
	if err != nil {
		t.Errorf("Unexpented RemoveComment error")
	}

	err = cl.RemoveComment(-1)

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}
