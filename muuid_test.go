package muuid

import "testing"

func TestOsxUUID(t *testing.T) {
	uuid, err := osxUUID()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuid)
}
