package pingdom

import "testing"

func TestCreateCheck(t *testing.T) {
	client := NewClient()
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestUpdateCheck(t *testing.T) {
	client := NewClient()
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.UpdateCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestDeleteCheck(t *testing.T) {
	client := NewClient()
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.DeleteCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
