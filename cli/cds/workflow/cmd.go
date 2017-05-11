package workflow

import "github.com/spf13/cobra"
import "github.com/ovh/cds/sdk"
import "fmt"

var (
	rootCmd = &cobra.Command{
		Use:   "workflow",
		Short: "cds workflow",
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List workflow on current project",
		Long:  "List workflow on current project.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				sdk.Exit("Wrong usage: %s\n", cmd.Short)
			}
			ws, err := sdk.WorkflowList(args[0])
			if err != nil {
				sdk.Exit("Error: %s\n", err)
			}
			sdk.Output("yaml", ws, fmt.Printf)
		},
	}

	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show a workflow on current project",
		Long:  "Show a  workflow on current project.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				sdk.Exit("Wrong usage: %s\n", cmd.Short)
			}
			ws, err := sdk.WorkflowGet(args[0], args[1])
			if err != nil {
				sdk.Exit("Error: %s\n", err)
			}
			sdk.Output("json", ws, fmt.Printf)
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
}

//Cmd returns the root command
func Cmd() *cobra.Command {
	return rootCmd
}
