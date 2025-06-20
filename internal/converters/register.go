package converters

import(
    "fmt"
)

var converters []Converter

func RegisterConverter(c Converter) {
    converters = append(converters, c)
}


func GetConverters() []Converter {

    fmt.Println("initialized GetConverters")
    return converters
}



