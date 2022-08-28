/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func fetchApp(app string, wg *sync.WaitGroup) {
	r := strings.Split(app, "->")
	app = r[0]

	if len(r) == 1 {
		r = append(r, "200")
	}

	expectedStatus, _ := strconv.Atoi(r[1])
	res, _ := http.Get(app)

	if expectedStatus != res.StatusCode {
		log.Printf("expected %d in %s, got %d", expectedStatus, app, res.StatusCode)
	} else {
		log.Printf("%s is ok", app)
	}

	wg.Done()
}

func fetchApps(apps []string) {
	wg := &sync.WaitGroup{}

	for range time.Tick(time.Second * 5) {
		for _, app := range apps {
			wg.Add(1)
			go fetchApp(app, wg)

			wg.Wait()
		}
	}

}

// wakeupCmd represents the wakeup command
var wakeupCmd = &cobra.Command{
	Use:   "wakeup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apps, _ := cmd.Flags().GetStringArray("apps")
		file, _ := cmd.Flags().GetString("file")

		if len(apps) > 0 {
			fetchApps(apps)
		}

		if file != "" {
			f, err := os.ReadFile(file)

			if err != nil {
				log.Fatal(err)
			}

			apps := strings.Split(string(f), "\n")

			fetchApps(apps)
		}

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		<-ch

		log.Println("shutdown...")
	},
}

func init() {
	rootCmd.AddCommand(wakeupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	wakeupCmd.PersistentFlags().StringArrayP("apps", "a", []string{}, "List of apps")
	wakeupCmd.PersistentFlags().StringP("file", "f", "", "File of apps")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wakeupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
