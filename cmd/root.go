package cmd


import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "Zonverter",
    Short: "A cli tool which will convert any file to any extension ",
    Long: `A cli tool which will convert any file to any extension `,


}


func Execute()error {
    return rootCmd.Execute()
}


func init() {
    rootCmd.AddCommand(hellocmd)
    rootCmd.AddCommand(convert)

}
