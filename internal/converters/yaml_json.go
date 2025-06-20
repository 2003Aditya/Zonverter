package converters

import (
	"encoding/json"
    "gopkg.in/yaml.v3"
	"fmt"
	"os"
	"path/filepath"
)

type JSONYAMLConverter struct{}

func (j JSONYAMLConverter) CanConvert(inputext, outputext string) bool {
    fmt.Println("Initialized CanConvert")
    return (inputext == ".json" && outputext == ".yaml") ||
    (inputext == ".yaml" && outputext == ".json")
}


func (j JSONYAMLConverter) Convert(inputPath, outputPath string) error {

    //Read the file
    file , err := os.ReadFile(inputPath)
    if err != nil {
        return err
    }


    inputExt := filepath.Ext(inputPath)
    fmt.Println(inputExt)

    outputExt := filepath.Ext(outputPath)
    fmt.Println(outputExt)

    //parsed will store the Unmarshal data
    var parsed any

    if inputExt == ".json" {
        err := json.Unmarshal(file, &parsed )
        fmt.Println(err)
        if err != nil {
            return err
        }
    } else if inputExt == ".yaml" {
        err := yaml.Unmarshal(file, &parsed)
        if err != nil {
            return err
        }
    }

    if outputExt == ".json" {
        json, err := json.MarshalIndent(parsed, "", " ")
        if err != nil {
            return err
        }

        err = os.WriteFile(outputPath, json, 0644)
        if err != nil {
            return err
        }


    } else if outputExt == ".yaml" {
        yaml, err := yaml.Marshal(parsed)
        if err != nil {
            return err
        }

        err = os.WriteFile(outputPath, yaml, 0644)
        if err != nil {
            return err
        }
    }

    fmt.Println("Initialized Convert")



    return nil
}

func init() {
    RegisterConverter(JSONYAMLConverter{})
}

