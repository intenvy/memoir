package key

const (
	SelectorChar = '*'
)

type Key string

func New(pattern string, converter Converter, validator Validator) *Key {
	if err := validator.Validate(pattern); err != nil {
		panic(err)
	}
	output := Key(converter.Convert(pattern))
	return &output
}

func (k *Key) Size() int {
	return len(string(*k))
}

func (k *Key) IsSelector() bool {
	return rune((*k)[k.Size()-1]) == SelectorChar
}

func (k *Key) IsRaw() bool {
	return !k.IsSelector()
}
