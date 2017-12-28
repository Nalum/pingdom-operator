package pingdomclient

import "testing"

func TestNewHTTPCheck(t *testing.T) {
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	if check.GetType() != CheckTypeHTTP {
		t.Fail()
		t.Log("Failed checking the Type of the Check")
	}
}

func TestNewHTTPCheckWithAuth(t *testing.T) {
	check, err := NewHTTPCheck("testing", "https://jim:car@this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	if check.GetType() != CheckTypeHTTP {
		t.Fail()
		t.Log("Failed checking the Type of the Check")
	}
}

func TestNewHTTPCheckWithPort(t *testing.T) {
	check, err := NewHTTPCheck("testing", "https://this.is:8080/a/test")

	if err != nil {
		t.Error(err)
	}

	if check.GetType() != CheckTypeHTTP {
		t.Fail()
		t.Log("Failed checking the Type of the Check")
	}
}

func TestHTTPCheckSetData(t *testing.T) {
	check, _ := NewHTTPCheck("testing", "https://this.is/a/test")

	err := check.SetData(map[string]interface{}{})

	if err != nil {
		t.Fail()
		t.Log(err)
	}
}

func TestHTTPCheckBadPort(t *testing.T) {
	_, err := NewHTTPCheck("testing", "https://this:99999999.is/a/test")

	if err == nil {
		t.Fail()
		t.Log("Error Expected but none given")
	}
}

func TestHTTPCheckBadURL(t *testing.T) {
	_, err := NewHTTPCheck("testing", "-https://this:99999999.is/a/test")

	if err == nil {
		t.Fail()
		t.Log("Error Expected but none given")
	}
}
