package muuid

import "testing"

func TestOsxUUID(t *testing.T) {
	uuid := UUID()
	if uuid == "" {
		t.Fatal("got empty uuid string")
	}
	t.Log(uuid)
}
