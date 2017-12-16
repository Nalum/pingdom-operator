package pingdom

import (
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	check := Check{
		Name: "Testing",
		URL:  "http://test.com/do-it",
	}

	err := check.Create()

	if err != nil {
		t.Fail()
		log.Println("Error: ", err)
	}
}

func TestUpdate(t *testing.T) {
	check := Check{
		Name: "Testing",
		URL:  "http://test.com/do-it",
	}

	err := check.Update()

	if err != nil {
		t.Fail()
		log.Println("Error: ", err)
	}
}
