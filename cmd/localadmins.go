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

var localadminsCommand = &cobra.Command{
	Use:    "localadmins [flags] <445.open>",
	Short:  "query members of local admins group",
	Long:   `Local administrators group on each queried for members`,
	Args:   cobra.ExactArgs(1),
	PreRun: setupSession,
	Run:    localAdmins,
}

func init() {
	rootCmd.AddCommand(localadminsCommand)
}

func localAdmins(cmd *cobra.Command, args []string) {
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
		go makeEnumWorker(ctx, winChan, &wg)
	}

	start := time.Now()

Scan:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break Scan
		default:
			winline := scanner.Text()
			//logger.Log.Debug("  trying host =>", winline)
			//>>>could check if IP is in right format.
			//winname, err := util.FormatTargets(winline)
			//if err != nil {
			//	logger.Log.Debug("[!] %q - %v", winline, err.Error())
			//	continue
			//}
			//logger.Log.Debug("  trying host =>", winname)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			//winChan <- winname
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

	// result, err := kSession.TestUsername(usernamelist)
	// if result {
	// 	fmt.Printf("[+] %v exists!\n", usernamelist)
	// }
	// if err != nil {
	// 	fmt.Println("erro!")
	// 	fmt.Printf(err.Error())
	// }
	// fmt.Println("Done!")
}
