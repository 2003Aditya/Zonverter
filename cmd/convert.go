package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hellocmd = &cobra.Command{
    Use:  "hello",
    Short: "A brief description of your command",
    Long: `A description... `,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("first command")
    },
}

func init() {
    rootCmd.AddCommand(hellocmd)
}
