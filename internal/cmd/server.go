package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	dhttp "github.com/turbolytics/dispatcher/internal/http"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the dispatcher",
	Long: `Starts the dispatcher and begins processing tasks. 
You can provide additional flags and arguments to customize its behavior.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dispatcher server starting on port :8080")
		s, err := dhttp.NewServer()
		if err != nil {
			panic(err)
		}

		r := dhttp.NewRoutes(s)
		http.ListenAndServe(":8080", r)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// You can define flags and configuration settings here.
	// e.g., serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
