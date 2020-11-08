package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "beachrpcgo",
	Short: "A tool to enumerate various ",
	Long: `This tool is
It is designed.
Warn accounts `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "FQ domain name (e.g. contoso.com or localhost)")
  rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Usernameee")
  rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")

//rootCmd.PersistentFlags().StringVar(&domainController, "dc", "", "Tlo")
//rootCmd.PersistentFlags().StringVarP(&logFileName, "output", "o", "", "File to write logs to. Optional.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, " ")
//rootCmd.PersistentFlags().BoolVar(&safe, "safe", false, "Safe mode. Will abort if -- . Default: FALSE")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 10, "Threads to use")
	rootCmd.PersistentFlags().IntVarP(&delay, "delay", "", 0, "Delay in ms between each attempt- single thread if set")
	rootCmd.PersistentFlags().StringVarP(&logZout, "out", "o", "", "Results csv out. Optional.")


	if delay != 0 {
		threads = 1
	}

}
