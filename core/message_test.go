package core

import "testing"

func TestPopulateTemplate(t *testing.T) {
	msg := Message{
		Body: "Hello, This is me.",
		Name: "Louise",
	}

	actual, err := PopulatTemplate(msg)

	if err != nil {
		t.Error(err)
	}

	t.Error(actual)
}
