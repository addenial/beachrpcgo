package cmd

import (
	"context"
	"os"
	// "../session"
	"../util"
	"github.com/spf13/cobra"
	"unicode/utf16"
	"golang.org/x/crypto/md4"
	"fmt"
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
	ntlm    				 string

	// Used for multithreading
	ctx, cancel = context.WithCancel(context.Background())
	counter     int32
	successes   int32
)




func utf16le(val string) []byte {
	var v []byte
	for _, r := range val {
		if utf16.IsSurrogate(r) {
			r1, r2 := utf16.EncodeRune(r)
			v = append(v, byte(r1), byte(r1>>8))
			v = append(v, byte(r2), byte(r2>>8))
		} else {
			v = append(v, byte(r), byte(r>>8))
		}
	}
	return v
}

func ntlmHash(password string) (hash [16]byte) {
	h := md4.New()
	h.Write(utf16le(password))
	h.Sum(hash[:0])
	return
}


func setupSession(cmd *cobra.Command, args []string) {
	logger = util.NewLogger(verbose, logFileName)
	if domain == "" {
		logger.Log.Error("No domain specified.. specify FQDN..")
		os.Exit(1)
	}


	//logic here -- len check, determine if plaintext or NT Hash was provided
  pwlen := len(password)

	ntlm = fmt.Sprintf("%x", ntlmHash(password))
  logger.Log.Debug("--pwnlen=>", pwlen)

  if pwlen == 32 {
    logger.Log.Debug("--len=32. ntlm provided. no need to hash.")
		ntlm = password
			} else {}




	logger.Log.Info("Using ->($$>(:0 --> ")
	logger.Log.Info("Domain:  ", domain)
	logger.Log.Info("User:    ", username)
	logger.Log.Info("Pass:    ", password)
  logger.Log.Info("NT Hash: ", ntlm)



	if delay != 0 {
		logger.Log.Infof("Delay set. Using single thread and delaying %dms between attempts\n", delay)
	}
}
