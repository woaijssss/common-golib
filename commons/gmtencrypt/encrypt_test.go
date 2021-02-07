package gmtencrypt

import "testing"

func TestDecrypt(t *testing.T) {

	if got := Decrypt("0766cc51a83b2de2113e762976cfaee8"); got != "" {
		t.Errorf("Decrypt() = %v, want %v", got, "")
	}
}
