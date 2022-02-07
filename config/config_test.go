package config

import "testing"

func TestSetConfig(t *testing.T) {
	SetConfig("./dev.yaml")
	t.Logf("%+v\n", Config.Get("notify"))
}
