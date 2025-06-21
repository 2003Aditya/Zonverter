package converters

import (
	"os"
	"testing"

	"github.com/2003Aditya/Zonverter/internal/converters"
)

func TestYAMLToJSON(t *testing.T) {

    input := `{"name": "Aditya"}`
    os.WriteFile("input.json",[]byte(input), 0644)

    converter := converters.JSONYAMLConverter{}
    err := converter.Convert("input.json", "output.json")
    if err != nil {
        t.Fatalf("Conversion failed:%v", err)
    }

    out, _ := os.ReadFile("output.yaml")
    if len(out) == 0 {
        t.Fatal("Output File is empty")
    }

    os.Remove("input.json")
    os.Remove("output.yaml")

}
