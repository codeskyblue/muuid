package muuid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var (
	ErrUuidNotFound = errors.New("uuid not found")
)

func UUID() string {
	var uuid string
	var err error

	switch runtime.GOOS {
	case "darwin":
		uuid, err = osxUUID()
	}
	if err != nil || uuid == "" {
		uuid = defaultUuid()
	}
	return uuid
}

func osxUUID() (string, error) {
	c := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	output, err := c.Output()
	if err != nil {
		return "", err
	}
	pattern := regexp.MustCompile(`IOPlatformUUID" = "(.*?)"`)
	ss := pattern.FindStringSubmatch(string(output))
	if len(ss) == 0 {
		return "", ErrUuidNotFound
	}
	return ss[1], nil
}

func linuxUUID() (string, error) {
	data, err := ioutil.ReadFile("/var/lib/dbus/machine-id")
	if err != nil {
		return "", ErrUuidNotFound
	}
	return "", nil
}

func defaultUuid() string {
	filePath := ".muid"
	id := ""
	data, err := ioutil.ReadFile(filePath)
	if err != nil || strings.TrimSpace(string(data)) == "" {
		id = fmt.Sprintf("%s", uuid.NewV4())
		ioutil.WriteFile(filePath, []byte(id), 0644)
	} else {
		return strings.TrimSpace(string(data))
	}
	return id
}
