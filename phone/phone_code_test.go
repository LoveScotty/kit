package phone

import "testing"

func TestGetPhoneCode(t *testing.T) {
	t.Log(GetPhoneCode("8613800138000"))
	t.Log(GetPhoneCode("+39339990   7440"))
}
