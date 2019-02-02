package hotreload

import (
	"encoding/json"
	"github.com/spiral/roadrunner/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockCfg struct{ cfg string }

func (cfg *mockCfg) Get(name string) service.Config  { return nil }
func (cfg *mockCfg) Unmarshal(out interface{}) error { return json.Unmarshal([]byte(cfg.cfg), out) }

func Test_Config_Hydrate(t *testing.T) {
	cfg := &mockCfg{`{"enable": true,"path":".","files":"*.php","tick":500}`}
	c := &Config{}

	assert.NoError(t, c.Hydrate(cfg))
}

func Test_Config_Hydrate_Default(t *testing.T) {
	cfg := &mockCfg{`{}`}
	c := &Config{}

	assert.NoError(t, c.Hydrate(cfg))

	assert.Equal(t, false, c.Enable)
	assert.Equal(t, ".", c.Path)
	assert.Equal(t, "*.php", c.Files)

	duration := time.Duration(500)
	assert.Equal(t, &duration, c.Tick)
}

func Test_Config_Hydrate_eeror(t *testing.T) {
	cfg := &mockCfg{`{"enable": true,"path":.,"files":"*.php","tick":500}`}
	c := &Config{}

	assert.Error(t, c.Hydrate(cfg))
}
