# beachrpcgo
Reconnaissance tool designed to assist with discovery of potential foothold hosts(beachhead hosts). Discovery is performed via net RPC calls over SMB (TCP445) protocol.

Tool will perform rapid discovery of local group memberships across Windows AD-joined systems. Requires valid active directory credentials- userland.

# usage

```
Usage:
  beachrpcgo [command]

Available Commands:
  help        Help about any command
  localadmins query members of local admins group
  rdpmembers  query members of remote desktop users group
  version     Display version info and quit
```

```
Mandatory flags:
  -d, --domain string     FQ domain name (e.g. contoso.com or localhost)
  -u, --username string   Username
  -p, --password string   Password plaintext or NT hash

Optional flags:
      --delay int         Delay in ms between each attempt- single thread if set
  -t, --threads int       Threads to use (default 10)
  -o, --out string        Save results out to csv file (e.g. out.csv)
  -h, --help              help for beachrpcgo
  -v, --verbose

Use "beachrpcgo [command] --help" for more information about a command.  
```


# examples:
# localadmins - local administrators group members
## targets from file (IP or hostname)
```
go run main.go localadmins -d localhost -u User1 -p Password123 445.open
```

## single target from stdin
```
echo 10.0.0.104 | go run main.go localadmins -d localhost -u User1 -p Password123 -
```

## save output to csv
```
go run main.go localadmins -d localhost -u User1 -p Password123 445.open -o outla.csv
```

# rdpmembers - remote desktop users group
```
go run main.go rdpmembers -d localhost -u User2 -p Password123 -o outrdp.csv -v 445.open
```
`cat outrdp.csv | grep Domain\ Users`



# compile instructions
to download golang environment in debian linux `sudo apt install golang`

compile without including debugging (recuse binary size by 30%~~ish)
`go build -ldflags "-w -s" -o beachrpcgo main.go`
`./beachrpcgo -h`

to specify binary architecture when compiling
`GOOS="linux" GOARCH="amd64" go build main.go`

beachrpcgo has a dependency on `net` (tool for administration of Samba and remote CIFS servers) and `pth-net`. tested on kali linux.
To download the patched pth kali tools which use password hashes as authentication input `apt install passing-the-hash -y`

# reference
versja edukacyjna golang
goroutine lekkie nici przez środowisko uruchomieniowe Go
shoutouts to ropnop for amaze code github.com/ropnop/kerbrute/
