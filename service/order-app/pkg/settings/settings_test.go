package settings

import "testing"

func TestLoadSetting(t *testing.T) {
	_, err := Load("")
	if err != nil {
		t.Error(err)
	}
}
