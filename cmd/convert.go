package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/2003Aditya/Zonverter/internal"
	"github.com/2003Aditya/Zonverter/pkg/utils"
	"github.com/spf13/cobra"
)


var convert = &cobra.Command {
    Use : "convert",
    Short : "converter for the file",
    Long : `A long description`,

    Run: func(cmd *cobra.Command, args []string) {

        input, err := cmd.Flags().GetString("input")
        if err != nil {
            log.Fatal("Error getting input flag", err)
        }

        output, err := cmd.Flags().GetString("output")
        if err != nil {
            log.Fatal("Error getting output file", err)
        }

        // force, err := cmd.Flags().GetString("force")
        // if err != nil {
        //     log.Fatal("Error getting force flag", err)
        // }

        fmt.Println("Input:", input)
        fmt.Println("Output:", output)
        // fmt.Println("Force:", force)

        checkInput := utils.CheckFile(input)
        checkOutput := utils.CheckFile(output)

        fmt.Println(checkInput)
        fmt.Println(checkOutput)

        inputExt := filepath.Ext(input)
        outputExt := filepath.Ext(output)

        internal.Dispatcher(inputExt, outputExt, input, output)



    },
}

func init() {
    rootCmd.AddCommand(convert)

    convert.Flags().String("input","", "Path to input file")
    convert.Flags().String("output","", "Path to output file")
    // convert.Flags().String("force","", "overwrite if output exists")

    convert.MarkFlagRequired("input")
    convert.MarkFlagRequired("output")
}
