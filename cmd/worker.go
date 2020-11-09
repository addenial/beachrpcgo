package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"sync/atomic"
	// "log"
	"bufio"
	"log"
	"strings"
	"os"
)

func makeEnumWorker(ctx context.Context, targets <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case target, ok := <-targets:
			if !ok {
				return
			}
			testTargetLA(ctx, target)
		}
	}
}

func testTargetLA(ctx context.Context, target string) {
	atomic.AddInt32(&counter, 1)


	logger.Log.Debug("* testing host =>", target)

	//userpass := "user2%Password123"
	//userpass := "user2%aad3b435b51404eeaad3b435b51404ee:NTLM
	userpass := fmt.Sprintf("%v%%aad3b435b51404eeaad3b435b51404ee:%v", username, ntlm)

	cmdZ := exec.Command("pth-net", "rpc", "group", "members", "administrators", "-U", userpass, "-S", target)
	outZ, errZ := cmdZ.CombinedOutput()
	if errZ != nil {
		logger.Log.Debugf("  unresponsive... %s", target)
	}
	//fmt.Printf("combined out:\n%s\n", string(outZ))
	//logger.Log.Noticef("[+] VALID USERNAME:\t %s", runmeplz)

	//logger.Log.Noticef("[+] VALID USERNAME:\t %s", logZout)

	//if logZout != "" {
		//outputFile, err := os.Create(logZout)
		//if err != nil {
			//panic(err)
		//}
	//} else {
	//}



	output := string(outZ)
	scanner := bufio.NewScanner(strings.NewReader(output))


	for scanner.Scan() {
		res1 := scanner.Text()
		newline := fmt.Sprintf("%s,%s", target, strings.Replace(res1, "\\", ",", 1 ))
		//fmt.Println(newline)
		logger.Log.Notice(newline)

		// if logging out to csv file
		if logZout != "" {
			f, err := os.OpenFile(logZout,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(newline + "\n"); err != nil {
				log.Println(err)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	//usernamefull := fmt.Sprintf("%v@%v", username, domain)

	//valid, err := kSession.TestUsername(username)
	//if valid {
	//	atomic.AddInt32(&successes, 1)
	//	logger.Log.Noticef("[+] VALID USERNAME:\t %s", usernamefull)
	//} else if err != nil {
	// This is to determine if the error is "okay" or if we should abort everything
	//	ok, errorString := kSession.HandleKerbError(err)
	//	if !ok {
	//		logger.Log.Errorf("[!] %v - %v", usernamefull, errorString)
	//		cancel()
	//	} else {
	//		logger.Log.Debugf("[!] %v - %v", usernamefull, errorString)
	//	}
	//} else {
	//	logger.Log.Debug("[!] Unknown behavior - %v", usernamefull)
	//}
}






func makeEnumWorker2(ctx context.Context, targets <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case target, ok := <-targets:
			if !ok {
				return
			}
			testTargetRDP(ctx, target)
		}
	}
}

func testTargetRDP(ctx context.Context, target string) {
	atomic.AddInt32(&counter, 1)

	userpass := fmt.Sprintf("%v%%aad3b435b51404eeaad3b435b51404ee:%v", username, ntlm)

	logger.Log.Debug("* testing host =>", target)
	cmdZ := exec.Command("pth-net", "rpc", "group", "members", "Remote Desktop Users", "-U", userpass, "-S", target)
	outZ, errZ := cmdZ.CombinedOutput()
	if errZ != nil {
		logger.Log.Debugf("  unresponsive... %s", target)
	}
	output := string(outZ)
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		res1 := scanner.Text()
		newline := fmt.Sprintf("%s,%s", target, strings.Replace(res1, "\\", ",", 1 ))
		//fmt.Println(newline)
		logger.Log.Notice(newline)

		// if logging out to csv file
		if logZout != "" {
			f, err := os.OpenFile(logZout,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(newline + "\n"); err != nil {
				log.Println(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
