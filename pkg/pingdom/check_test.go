package pingdom

import "testing"

func TestNewCheck(t *testing.T) {
	c := newCheck("Check Name", "checkhost.com", CheckTypeHTTP)

	if c.Name != "Check Name" {
		t.Fail()
	}

	if c.Host != "checkhost.com" {
		t.Fail()
	}

	if c.Type != CheckTypeHTTP {
		t.Fail()
	}
}

func TestGetType(t *testing.T) {
	c := newCheck("Check Name", "checkhost.com", CheckTypeHTTP)

	if c.GetType() != CheckTypeHTTP {
		t.Fail()
	}
}
