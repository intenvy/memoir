package key

// handles the conversion of keys, queries and key
// to satisfiable formats for the repository
type Converter interface {
	Convert(key string) string
}

// implements Converter
// converts key in steps, in a pipeline manner
type ConverterPipeline struct {
	Hooks []func(string) string
}

var _ Converter = (*ConverterPipeline)(nil)

func NewConverterPipeline() *ConverterPipeline {
	return &ConverterPipeline{Hooks: make([]func(string) string, 0)}
}

func (p *ConverterPipeline) AddHook(hooks ...func(string) string) *ConverterPipeline {
	p.Hooks = append(p.Hooks, hooks...)
	return p
}

func (p *ConverterPipeline) Convert(pattern string) string {
	for _, hook := range p.Hooks {
		pattern = hook(pattern)
	}
	return pattern
}
