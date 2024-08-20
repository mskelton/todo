package root

import (
	"fmt"
	"os"
	"path"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/todo/pkg/cmd/project"
	"github.com/mskelton/todo/pkg/cmd/today"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todoist CLI",
	Long: heredoc.Doc(`
    Welcome to the Todoist CLI!

    Running the todo command will display the home view of the task list.
    You can configure the home view in your Todoist settings, which will be
    respecting when running this command.
  `),
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return today.TodayCmd.RunE(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/todo/config.json)")

	rootCmd.AddCommand(today.TodayCmd)
	rootCmd.AddCommand(project.ProjectCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(path.Join(home, ".config", "todo"))
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; use defaults
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
