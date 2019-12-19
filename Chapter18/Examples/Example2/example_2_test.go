package command_injection
import "testing"
func TestListFiles(t *testing.T) {
	out, err := listFiles(" .; cat /etc/hosts")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}
