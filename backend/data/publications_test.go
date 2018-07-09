package data

import "testing"

var testPublications = []Publication{
	{
		Namepub:    "Backend",
		Time:       "21:00",
		Publcation: "Run backend test...",
	},
	{
		Namepub:    "Frontend",
		Time:       "00:00",
		Publcation: "Run frontend test...",
	},
}

func TestAddPublication(t *testing.T) {
	cl := NewPublicationlist()

	cl.AddPublication(0, testPublications[0])

	if cl.GetPublications()[0] != testPublications[0] {
		t.Errorf("AddPublication is not working")
	}
}

func TestEditPublication(t *testing.T) {
	cl := NewPublicationlist()

	cl.AddPublication(0, testPublications[0])

	err := cl.EditPublication(0, testPublications[1], 0)

	if cl.GetPublications()[0] != testPublications[1] {
		t.Errorf("EditPublication is not working")
	}
	if err != nil {
		t.Errorf("Unexpented EditPublication error")
	}

	err = cl.EditPublication(0, testPublications[1], -1)

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}

func TestDeletePublication(t *testing.T) {
	cl := NewPublicationlist()
	cl.AddPublication(0, testPublications[0])
	cl.AddPublication(0, testPublications[1])

	err := cl.RemovePublication(0, 0)

	if cl.GetPublications()[0] != testPublications[1] {
		t.Errorf("RemovePublication is not working")
	}
	if err != nil {
		t.Errorf("Unexpented RemovePublication error")
	}

	err = cl.RemovePublication(0, -1)

	if err == nil {
		t.Errorf("Not Handled out of range error")
	}
}
