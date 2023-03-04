package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/topisenpai/gobin/internal/ezhttp"
)

func NewVersionCmd(parent *cobra.Command, version string) {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of the gobin cli",
		Long: `Prints the version of the gobin cli. For example:

gobin version

Go Version: go1.18.3
Version: dev
Commit: b1fd421
Build Time: Mon Jan  1 00:00:00 0001
OS/Arch: windows/amd64

Go Version: go1.19
Version: dev
Commit: b1fd421
Build Time: Mon Jan  1 00:00:00 0001
OS/Arch: windows/amd64`,
		Run: func(cmd *cobra.Command, args []string) {
			server := viper.GetString("server")
			cmd.Println(version)

			if server != "" {
				rs, err := ezhttp.Get("/version")
				if err != nil {
					cmd.PrintErrln("Failed to get server version:", err)
					return
				}
				defer rs.Body.Close()

				data, err := io.ReadAll(rs.Body)
				if err != nil {
					cmd.PrintErrln("Failed to read server version:", err)
					return
				}
				cmd.Printf("Server: %s\n%s\n", server, data)
			}
		},
	}

	parent.AddCommand(cmd)

	cmd.Flags().StringP("server", "s", "", "Gobin server address")

	viper.BindPFlag("server", cmd.PersistentFlags().Lookup("server"))
}
