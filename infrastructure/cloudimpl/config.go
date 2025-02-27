package cloudimpl

import "github.com/opensourceways/xihe-inference-evaluate/infrastructure/config"

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

type Config struct {
	CRD          config.CRDConfig `json:"crd"              required:"true"`
	RPCEndpoint  string           `json:"rpc_endpiont"     required:"true"`
	JupyterToken string           `json:"jupyter_token"    required:"true"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.CRD,
	}
}

func (cfg *Config) SetDefault() {
	items := cfg.configItems()

	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) Validate() error {
	items := cfg.configItems()

	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
