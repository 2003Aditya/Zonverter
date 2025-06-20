package internal

import (
	"fmt"

	"github.com/2003Aditya/Zonverter/internal/converters"
)


func Dispatcher(inputExt, outputExt, inputPath, outputPath string) error {
    fmt.Println("initialized dispatcher")
    for _, converter := range converters.GetConverters() {
        if converter.CanConvert(inputExt, outputExt) {
            return converter.Convert(inputPath, outputPath)
        }
    }
    return fmt.Errorf("no converter was found for %s - %s", inputExt, outputExt)
}
