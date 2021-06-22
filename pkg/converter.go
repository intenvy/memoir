package pkg

// handles the conversion of keys, queries and patterns
// to satisfiable formats for the repository
type KeyConverter interface {
	Convert(key string) string
}

// implements KeyConverter
// converts key in steps, in a pipeline manner
type KeyPipeline struct {
	hooks []func(string) string
}

var _ KeyConverter = (*KeyPipeline)(nil)

func NewKeyPipeLine() *KeyPipeline {
	return &KeyPipeline{hooks: make([]func(string) string, 0)}
}

func (p *KeyPipeline) WithHook(hooks ...func(string) string) *KeyPipeline {
	p.hooks = append(p.hooks, hooks...)
	return p
}

func (p *KeyPipeline) Convert(key string) string {
	for _, hook := range p.hooks {
		key = hook(key)
	}
	return key
}
