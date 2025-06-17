package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)


var convert = &cobra.Command {
    Use : "convert",
    Short : "converter for the file",
    Long : `A long description`,

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("converter")
    },
}
