package dashboard

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func timeSinceModified(filename string) (time.Duration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	modifiedTime := fileInfo.ModTime()
	timeSinceModified := time.Since(modifiedTime)
	return timeSinceModified, nil
}

func isModifiedOver2MinutesAgo(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}

	modifiedTime := fileInfo.ModTime()
	timeSinceModified := time.Since(modifiedTime)
	if timeSinceModified > (2 * time.Minute) {
		return true, nil
	}
	return false, nil
}

func mrefdUptime(pidFile, checkFile string) (time.Duration, error) {
	uptime, err := timeSinceModified(pidFile)
	if err != nil {
		return 0, err
	}
	stale, err := isModifiedOver2MinutesAgo(checkFile)
	if err != nil {
		return 0, err
	}
	if stale {
		return uptime, fmt.Errorf("check file is older than two minutes")
	}
	return uptime, nil
}

// maskIP partially hides an IP address so only the first two
// octets or segments are returned. For IPv4 the result will
// look like "X.X.xx.xx" and for IPv6 "X:X:xx:xx...".
func maskIP(ip string) string {
	if ip == "" {
		return ""
	}
	if strings.Contains(ip, ".") {
		parts := strings.Split(ip, ".")
		if len(parts) >= 2 {
			return fmt.Sprintf("%s.%s.xx.xx", parts[0], parts[1])
		}
	}
	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")
		if len(parts) >= 2 {
			masked := make([]string, len(parts)-2)
			for i := range masked {
				masked[i] = "xx"
			}
			return strings.Join(append([]string{parts[0], parts[1]}, masked...), ":")
		}
	}
	return ip
}

// if needed to test the function
/*func main() {
	pidFile := "/var/run/mrefd.pid"
	checkFile := "/var/log/mrefd.xml"
	uptime, err := mrefdUptime(pidFile, checkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("mrefd has been up for %s", uptime)
}
*/
