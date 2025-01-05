package conf

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init(context.Background(), "sim.toml")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(GlobalC)
}
