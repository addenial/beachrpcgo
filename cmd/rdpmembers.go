package cmd

import (
	"bufio"
	"os"
	"sync"
	"sync/atomic"
	"time"

	//"../util"
	"github.com/spf13/cobra"
)

var rdpmembersCommand = &cobra.Command{
	Use:    "rdpmembers [flags] <445.open>",
	Short:  "query members of remote desktop users group",
	Long:   `Queries and shows members of the Remote Desktop Users group`,
	Args:   cobra.ExactArgs(1),
	PreRun: setupSession,
	Run:    rdpmembers,
}

func init() {
	rootCmd.AddCommand(rdpmembersCommand)
}

func rdpmembers(cmd *cobra.Command, args []string) {
	//logger.Log.Infof("IP,host,group member")
	wintargets := args[0]
	winChan := make(chan string, threads)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(threads)

	var scanner *bufio.Scanner
	if wintargets != "-" {
		file, err := os.Open(wintargets)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for i := 0; i < threads; i++ {
		go makeEnumWorker2(ctx, winChan, &wg)
	}

	start := time.Now()

Scan:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break Scan
		default:
			winline := scanner.Text()
			time.Sleep(time.Duration(delay) * time.Millisecond)
			winChan <- winline
		}
	}
	close(winChan)
	wg.Wait()

	finalCount := atomic.LoadInt32(&counter)
	finalSuccess := atomic.LoadInt32(&successes)
	logger.Log.Infof("Done!  %d targetz selected (%d dee) in %.3f seconds", finalCount, finalSuccess, time.Since(start).Seconds())

	if err := scanner.Err(); err != nil {
		logger.Log.Error(err.Error())
	}

}
