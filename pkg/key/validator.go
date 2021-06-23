package key

// validates the Pattern of keys
type Validator interface {
	Validate(pattern string) error
}

// implements Validator
// validates keys in steps, in a pipeline manner
type ValidatorPipeline struct {
	Hooks []func(string) error
}

var _ Validator = (*ValidatorPipeline)(nil)

func NewValidatorPipeline() *ValidatorPipeline {
	return &ValidatorPipeline{Hooks: make([]func(string) error, 0)}
}

func (v *ValidatorPipeline) AddHook(hooks ...func(string) error) *ValidatorPipeline {
	v.Hooks = append(v.Hooks, hooks...)
	return v
}

func (v *ValidatorPipeline) Validate(pattern string) error {
	for _, hook := range v.Hooks {
		if err := hook(pattern); err != nil {
			return err
		}
	}
	return nil
}
