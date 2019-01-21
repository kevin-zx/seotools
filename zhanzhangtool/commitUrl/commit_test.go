package commitUrl

import "testing"

func TestCommit(t *testing.T) {
	_, err := Commit("vB7VZYiFiHOEu4S9", "https://www.infineon-autoeco.com", []string{"https://www.infineon-autoeco.com/KnowledgeBase/Detail/939"})
	if err != nil {
		t.Error(err)
	}
}
