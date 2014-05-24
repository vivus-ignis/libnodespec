package libnodespec

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// socketValue is a port number for tcp/udp or an absolute path to a unix socket file
func translateSocketInode(inode int64) (socketType string, socketValue string, err error) {
	procNetTcp, err := ioutil.ReadFile("/proc/net/tcp")
	if err != nil {
		return "", "", err
	}
	// procNetTcp6, err := ioutil.ReadFile("/proc/net/tcp6")
	// if err != nil {
	// 	return "", "", err
	// }

	// procNetUdp, err := ioutil.ReadFile("/proc/net/udp")
	// if err != nil {
	// 	return "", "", err
	// }
	// procNetUdp6, err := ioutil.ReadFile("/proc/net/udp6")
	// if err != nil {
	// 	return "", "", err
	// }

	for _, line := range strings.Split(string(procNetTcp), "\n") {
		fields := strings.Fields(line)
		if inodeField, err := strconv.ParseInt(fields[9], 10, 0); inodeField == inode && err == nil {
			portHex := strings.Split(fields[2], ":")[1]
			portDec, err := strconv.ParseInt(portHex, 16, 0)
			if err != nil {
				return "", "", err
			}

			return "tcp", string(portDec), nil
		}
	}

	return "", "", errors.New("Translation failed - inode unknown to OS")
}

func (spec SpecService) Run(defaults PlatformDefaults) (err error) {
	psDirs, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		return err
	}
	for _, processDir := range psDirs {
		psName, err := ioutil.ReadFile(path.Join(processDir, "cmdline"))
		if err != nil {
			return err
		}
		if len(psName) == 0 {
			continue
		}
		// stripping \0
		psNameLen := len(psName) - 1

		if string(psName[:psNameLen]) == spec.Name {
			processSockets := map[string]map[string]bool{
				"tcp":  map[string]bool{},
				"udp":  map[string]bool{},
				"unix": map[string]bool{},
			}

			if len(spec.Sockets) > 0 || len(spec.Ports) > 0 {
				fdFiles, err := filepath.Glob(path.Join(processDir, "fd", "[0-9]*"))
				if err != nil {
					return err
				}
				for _, fdSymlink := range fdFiles {
					if fdNum, err := strconv.ParseInt(fdSymlink, 10, 0); fdNum < 2 && err == nil {
						continue // skip stdin, stdout, stderr
					}

					fdSymlinkTo, err := os.Readlink(fdSymlink)
					if err != nil {
						continue
					}

					inodeRe := regexp.MustCompile(`socket:\[(0-9+)\]`)
					if strings.HasPrefix(fdSymlinkTo, "socket:") {
						socketInode, err := strconv.ParseInt(inodeRe.FindStringSubmatch(fdSymlinkTo)[1], 10, 0)
						if err != nil {
							continue // FIXME
						}
						socketType, socketValue, err := translateSocketInode(socketInode)
						if err == nil {
							if socketType == "unix" {
								processSockets["unix"][socketValue] = true
							} else if socketType == "tcp" {
								processSockets["tcp"][socketValue] = true
							} else if socketType == "udp" {
								processSockets["udp"][socketValue] = true
							}
						}
					}

				}

				for _, wantedUnixSocket := range spec.Sockets {
					if _, found := processSockets["unix"][wantedUnixSocket]; found {
						return errors.New("Process doesn't own socket " + wantedUnixSocket)
					}
				}
				for _, wantedTcpUdpSocket := range spec.Ports {
					family := strings.Split(wantedTcpUdpSocket, ":")[0]
					portNumber := strings.Split(wantedTcpUdpSocket, ":")[1]
					if _, found := processSockets[family][portNumber]; found {
						return errors.New("Process doesn't listen on port " + wantedTcpUdpSocket)
					}
				}

			}

			return nil
		}
	}

	return errors.New("No such process")
}
