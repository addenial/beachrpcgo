package cmd

import (
	"context"
	"os"

	// "../session"
	"../util"
	"github.com/spf13/cobra"
)

var (
	domain           string
  username         string
  password         string
	domainController string
	logFileName      string
	verbose          bool
	safe             bool
	delay            int
	threads          int
	stopOnSuccess    bool
	userAsPass       = false
	logger           util.Logger
	// kSession         session.KerbruteSession
	logZout					 string

	// Used for multithreading
	ctx, cancel = context.WithCancel(context.Background())
	counter     int32
	successes   int32
)

func setupSession(cmd *cobra.Command, args []string) {
	logger = util.NewLogger(verbose, logFileName)
	if domain == "" {
		logger.Log.Error("No domain specified.. specify FQDN..")
		os.Exit(1)
	}
	//var err error
	// kSession, err = session.NewKerbruteSession(domain, domainController, verbose, safe)
	//if err != nil {
	//	logger.Log.Error(err.Error())
	//	os.Exit(1)
	//}
	logger.Log.Info("Using ->($$>(:0 --> ")
	logger.Log.Info("Domain:  ", domain)
	logger.Log.Info("User:    ", username)
	logger.Log.Info("Pass:    ", password)
//logger.Log.Info("NT Hash: ", ntlm)


	//for _, v := range kSession.Kdcs {
	//	logger.Log.Infof("\t%s\n", v)
	//}
	if delay != 0 {
		logger.Log.Infof("Delay set. Using single thread and delaying %dms between attempts\n", delay)
	}
}
