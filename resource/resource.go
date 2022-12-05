// Package resource
// @author： Boice
// @createTime：2022/11/28 11:19
package resource

type (
	Resource interface {
		Logger() Logger
		Config() Config
	}

	resource struct {
		logger Logger
		config Config
	}
)

func (r *resource) Logger() Logger {
	return r.logger
}

func (r *resource) Config() Config {
	return r.config
}

func New(configPath string) Resource {
	conf := newConfig(configPath)
	log := newLogger(conf)
	return &resource{
		logger: log,
		config: conf,
	}
}
