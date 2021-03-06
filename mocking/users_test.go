package task

import (
	"testing"
)

func TestMyApp_FilterByKeyword(t *testing.T) {
	userFetcher := UserFetcher{}
	app := SimpleApp{&userFetcher}
	keywords := []string{"server", "net", "multimedia", "task"}
	users, err := app.FilterByKeyword(keywords)
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 3 {
		t.Errorf("Expected 3 users, got %d", len(users))
	}
}
