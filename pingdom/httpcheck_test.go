package pingdom

import "testing"

func TestNewHTTPCheck(t *testing.T) {
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Fail()
		t.Error(err)
	}

	if check.GetType() != CheckTypeHTTP {
		t.Fail()
		t.Log("Failed checking the Type of the Check")
	}
}
