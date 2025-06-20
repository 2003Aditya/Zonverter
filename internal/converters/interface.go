package converters

type Converter interface {
    CanConvert(inputext, outputext string) bool
    Convert(inputext, outputext string) error
}


