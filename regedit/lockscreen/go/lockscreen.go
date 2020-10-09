package main

import (
	"errors"
	"io"
	"log"
	"os/user"

	"golang.org/x/sys/windows/registry"
)

// func main() {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	d, err := getLockScreenPath()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(d)
// }

func getUserSID() (string, error) {
	uinstance, err := user.Current()
	if err != nil {
		return "", err
	}
	// https://pkg.go.dev/os/user#User
	// Uid is the Sid on windows
	return uinstance.Uid, err
}

// ErrNotSpotLight means currently windows spotlight is turned off
var ErrNotSpotLight = errors.New("Not Windows spotlight")

func getLockScreenRegKeySpotLight() (string, error) {
	// win.GetCurrentProcessId()
	sid, err := getUserSID()
	if err != nil {
		return "", err
	}
	ret := `SOFTWARE\Microsoft\Windows\CurrentVersion\Authentication\LogonUI\Creative\` + sid
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, ret, registry.READ)
	defer key.Close()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	windowsSpotlight, _, err := key.GetIntegerValue("RotatingLockScreenEnabled")
	if err != nil {
		return "", err
	}
	if windowsSpotlight == 1 {
		subKeys, err := key.ReadSubKeyNames(100)
		if (err != nil && err != io.EOF) || len(subKeys) == 0 {
			log.Println(err)
			log.Fatal(err, " Subkeys doesn't exist, possibly wrong place to look for lockscreen ", len(subKeys))
			return "", err
		}
		return ret + `\` + subKeys[len(subKeys)-1], nil
	}
	return "", ErrNotSpotLight
}

func getLockScreenPath() (string, error) {
	lockScreenRegKey, err := getLockScreenRegKeySpotLight()
	if err != nil {
		if err == ErrNotSpotLight {
			// TODO not windows spotlight get from somewhere else
			// check what registry changes
			log.Println("Not windows spotlight")
			return "", nil
		}
		log.Fatal(err)
		return "", err
	}
	// Spotlight
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, lockScreenRegKey, registry.READ)
	defer key.Close()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	p, _, err := key.GetStringValue("landscapeImage")
	if err != nil {
		return "", err
	}
	return p, nil
}
