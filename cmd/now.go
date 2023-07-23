/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
  "time"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Get the current unix timestamp",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: runCmd,
}

func runCmd(cmd *cobra.Command, args []string) {
  if timezone, _ := cmd.Flags().GetString("timezone"); timezone != "" {
    time.LoadLocation(timezone)
  }

  if tzoff, _ := cmd.Flags().GetInt("tzoff"); tzoff != 0 {
    time.FixedZone("tzoff", tzoff)
  }

  var now time.Time

  if local, _ := cmd.Flags().GetBool("local"); !local {
    now = time.Now().UTC()
  } else {
    now = time.Now()
  }

  if delta, _ := cmd.Flags().GetString("delta"); delta != "" {
    duration, err := time.ParseDuration(delta)
    if err != nil {
      fmt.Println(err)
      return;
    }
    now = now.Add(duration)
  }
  if human, _ := cmd.Flags().GetBool("human"); human {
    fmt.Println(now)
    return;
  }
  fmt.Println(now.Unix())
}

func init() {
	getCmd.AddCommand(nowCmd)

	nowCmd.Flags().BoolP("human", "u", false, "Human readable output")
	nowCmd.Flags().BoolP("millis", "m", false, "Use milliseconds")
	nowCmd.Flags().BoolP("local", "l", false, "Use local time instead of UTC")
	nowCmd.Flags().IntP("tzoff", "o", 0, "Timezone offset")
	nowCmd.Flags().StringP("timezone", "t", "", "Timezone")
  nowCmd.Flags().StringP("delta", "d", "", "Add/minus delta")
}
