package converters

import (
    "encoding/xml"
    "gopkg.in/yaml.v3"
    "fmt"
    "os"
    "path/filepath"
    "github.com/clbanning/mxj/v2"
)

type YAMLXMLConverter struct{}

func( j YAMLXMLConverter) CanConvert(inputext, outputext string) bool{

    fmt.Println("Initialized YAMLXMLConverter")
    return (inputext == ".yaml" && outputext == ".xml") ||
    (inputext == ".xml" && outputext == ".yaml")
}

func (j YAMLXMLConverter) Convert(inputPath, outputPath string) error {

    fmt.Println("initialized convert")

    file, err := os.ReadFile(inputPath)
    fmt.Println(err)
    if err != nil {
        return err
    }

    inputext := filepath.Ext(inputPath)
    fmt.Println(inputext)
    outputext := filepath.Ext(outputPath)
    fmt.Println("OutputExt:" ,outputext)

    var parsed any

    if inputext == ".yaml" {
        err := yaml.Unmarshal(file, &parsed)

        fmt.Println("error for inputExt Unmarshal:", err)
        if err != nil {
            return err
        }
    } else if inputext == ".xml" {
        err := xml.Unmarshal(file, &parsed)
        fmt.Println(err)
        if err != nil {
            return err
        }
    }

    if outputext == ".xml" {
        fmt.Println("inside slice")

        switch v := parsed.(type) {
        case map[string]interface{}:
            fmt.Println("xml is structured")
            mv := mxj.Map(v)
            xmlBytes, err := mv.XmlIndent("", " ")
            if err != nil {
                return err
            }
            return os.WriteFile(outputPath,xmlBytes, 0644 )

        case []interface{}:
            fmt.Println("xml is unstrc")
            mv := mxj.Map{"root":v}
            xmlBytes, err := mv.XmlIndent("", " ")
            if err != nil {
                return err
            }

            return os.WriteFile(outputPath, xmlBytes, 0644)

        }
        fmt.Println("done with xml")


    } else if outputext == ".yaml" {
        mv, err := mxj.NewMapXml(file)
        if err != nil {
            return err
        }
        // fmt.Println(xml)

        yaml, err := yaml.Marshal(mv)
        if err != nil {
            return err
        }

        err = os.WriteFile(outputPath, yaml, 0644)
        if err != nil {
            return err
        }
    }
    fmt.Println("done with json_xml")
    return nil

}

func init() {
    RegisterConverter(YAMLXMLConverter{})
}

