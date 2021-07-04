package mapper

import "github.com/mitchellh/mapstructure"

// Mapper mapper instance
type Mapper struct {
	weakInput bool
}

// NewMapper returns new mapper instance
func NewMapper(weakInput bool) *Mapper {
	return &Mapper{
		weakInput: weakInput,
	}
}

// Decode decodes data from origin struct to target
func (m *Mapper) Decode(origin interface{}, dest interface{}) error {
	decoder, err := m.getDecoder(dest)
	if err != nil {
		return err
	}

	decoder.Decode(origin)

	return nil
}

func (m *Mapper) getDecoder(dest interface{}) (*mapstructure.Decoder, error) {
	decoderConfig := m.getDecoderConfig(dest)

	mapper, err := mapstructure.NewDecoder(&decoderConfig)
	if err != nil {
		return &mapstructure.Decoder{}, err
	}

	return mapper, nil
}

func (m *Mapper) getDecoderConfig(dest interface{}) mapstructure.DecoderConfig {
	mapConfig := mapstructure.DecoderConfig{
		Result:           dest,
		WeaklyTypedInput: m.weakInput,
	}

	return mapConfig
}
