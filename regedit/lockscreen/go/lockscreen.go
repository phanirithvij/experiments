package main

import (
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os/user"

	"golang.org/x/sys/windows/registry"
)

func main() {
	log.SetFlags(0)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// d, err := getLockScreenPath()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(d)
	getSysLockScreenConfig()
}

func getSysLockScreenConfig() {
	// sid, err := getUserSID()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data, err := hex.DecodeString("9a19a1622d473e54073d5bd08883f92a1852")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	// S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen
	ret := `SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen`
	// ret := sid + `\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen`
	key, err := registry.OpenKey(registry.CURRENT_USER, ret, registry.READ)
	// key, err := registry.OpenKey(registry.USERS, ret, registry.READ)
	defer key.Close()
	if err != nil {
		log.Fatal(err)
	}
	subValues, err := key.ReadValueNames(100)
	for _, v := range subValues {
		_, b, err := key.GetValue(v, nil)
		if err != nil {
			log.Fatal(err)
		}

		if b == registry.SZ || b == registry.EXPAND_SZ {
			// data, _, _ := key.GetStringValue(v)
			// log.Println("string", data)
		} else if b == registry.BINARY {
			bin, _, _ := key.GetBinaryValue(v)
			log.Println(v, "binary", len(bin))
			data := []byte{}
			for i := 0; i < len(bin); i++ {
				if bin[i] != 0 {
					data = append(data, bin[i])
				}
			}
			log.Println(string(data), len(data))
		} else if b == registry.DWORD || b == registry.DWORD_BIG_ENDIAN {
			// data, _, _ := key.GetIntegerValue(v)
			// log.Println("integer", data)
		} else {
			// var buf []byte
			// buf = make([]byte, n)
			// n, b, err = key.GetValue(v, buf)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// log.Println("unknown type", buf)
		}

		// log.Println(key.GetStringValue())
	}
	if (err != nil && err != io.EOF) || len(subValues) == 0 {
		log.Fatal(err)
	}
}

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
