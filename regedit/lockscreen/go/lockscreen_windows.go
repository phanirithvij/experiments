package main

import (
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/spf13/afero"

	"golang.org/x/sys/windows/registry"
)

func main() {
	// log.SetFlags(0)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	d, err := getLockScreenPath()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d)
	// fmt.Println(d)
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
	key, err := registry.OpenKey(registry.CURRENT_USER, ret, registry.READ|registry.WOW64_64KEY)
	// key, err := registry.OpenKey(registry.USERS, ret, registry.READ|registry.WOW64_64KEY)
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
			// log.Println(data, len(data))
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
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, ret, registry.READ|registry.WOW64_64KEY)
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
	sid, err := getUserSID()
	if err != nil {
		log.Fatal(err)
	}
	lockScreenRegKey, err := getLockScreenRegKeySpotLight()
	if err != nil {
		if err == ErrNotSpotLight {
			// not windows spotlight could be one of picture/slideshow
			log.Println("Not windows spotlight")
			ret := `SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen`
			key, err := registry.OpenKey(registry.CURRENT_USER, ret, registry.READ|registry.WOW64_64KEY)
			slideshowEnabled, _, err := key.GetIntegerValue("SlideshowEnabled")
			if err != nil {
				log.Println(err)
				return "", nil
			}
			if slideshowEnabled == 1 {
				// check what registry changes
				// This is close to impossible as I have non lead
				log.Println("Slide show")
				ret = `SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen`
				key, err := registry.OpenKey(registry.CURRENT_USER, ret, registry.READ|registry.WOW64_64KEY)
				spath, _, err := key.GetStringValue("SlideshowDirectoryPath1")
				if err != nil {
					if err == registry.ErrNotExist {
						log.Println("No slideshow directories selected by user")
					} else {
						log.Println(err)
					}
				} else {
					// TODO decode SlideshowDirectoryPath1 etc.
					log.Println("Slideshow path encoded", spath)
					log.Fatal("Slideshow not IMPLEMENTED")
				}
			} else {
				// Picture
				// TODO decode OriginalFile_A
				// printing as a string shows some kind of pattern can try
				ret = `SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\` + sid + `\AnyoneRead\LockScreen`
				log.Println("Picture")
				key, err := registry.OpenKey(registry.LOCAL_MACHINE, ret, registry.READ|registry.WOW64_64KEY)
				s, _, err := key.GetStringValue("")
				if err != nil {
					log.Println(err)
				}
				// In the order there can be Letters L in A-Z
				// And they keys OriginalFile_{L} may or may not exist
				// For eg. OriginalFile_Z does not exist
				if len(s) == 0 {
					log.Fatal("No pictures selected by user")
				}
				log.Println("Order", s)
				var labelx byte = s[0]
				filename, err := getFakeLockScreenFilePath(labelx)
				if err != nil {
					if _, ok := err.(*os.PathError); ok {
						// labled file doesn't exist
						// might be one from C:\Windows\Web\Screen
						webScreenDir := `C:\Windows\Web\Screen`
						// eg. Z is C:\Windows\Web\Screen\img100.jpg
						// TODO Assuming Z -> 100, Y -> 101, X -> 102 etc.
						inum := 190 - labelx
						if inum >= 100 && inum <= 105 {
							str := strconv.FormatInt((int64)(inum), 10)
							filename = webScreenDir + `\img` + str + ".jpg"
							_, err := os.Stat(filename)
							if err != nil {
								// dirty trick? No. Windows is dirty
								// if jpg not found png
								filename = webScreenDir + `\img` + str + ".png"
								_, err := os.Stat(filename)
								if err != nil {
									return "", err
								}
							}
							return filename, nil
						}
						return "", errors.New("Couldn't get the image")
					}
					log.Fatal(err)
				}
				// getSysLockScreenConfig()
				// sysDataRecover()
				return filename, nil
			}
			return "", nil
		}
		log.Fatal(err)
		return "", err
	}
	// Spotlight
	log.Println("Spotlight")
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, lockScreenRegKey, registry.READ|registry.WOW64_64KEY)
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

func sysDataRecover() {
	sid, err := getUserSID()
	if err != nil {
		log.Fatal(err)
	}
	systemData := `C:\ProgramData\Microsoft\Windows\SystemData\` + sid + `\ReadOnly`
	base := afero.NewOsFs()
	folder, err := base.Open(systemData)
	defer folder.Close()
	if err != nil {
		log.Fatal(err)
	}
	afero.Walk(base, folder.Name(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// TODO file can be PNG with .jpg extension
		// Need to check file magic info or something
		if !info.IsDir() && strings.HasSuffix(info.Name(), "LockScreen.jpg") {
			log.Println(path, info.Size())
		}
		return nil
	})
}

func getFakeLockScreenFilePath(label byte) (string, error) {
	sid, err := getUserSID()
	if err != nil {
		return "", err
	}
	systemData := `C:\ProgramData\Microsoft\Windows\SystemData\` + sid + `\ReadOnly\LockScreen_` + string(label) + `\LockScreen.jpg`
	lockscrfile, err := os.Open(systemData)
	defer lockscrfile.Close()
	if err != nil {
		return "", err
	}
	return lockscrfile.Name(), nil
}
