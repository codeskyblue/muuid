package muuid

import "testing"

func TestOsxUUID(t *testing.T) {
	uuid, err := linuxUUID()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuid)
}
