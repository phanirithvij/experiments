// this transcodes transcodeimagecache to a string
// look at getwall.ps1
// it's contents are this

// # https://www.winhelponline.com/blog/find-current-wallpaper-file-path-windows-10/
// $TIC=(Get-ItemProperty 'HKCU:\Control Panel\Desktop' TranscodedImageCache -ErrorAction Stop).TranscodedImageCache
// [System.Text.Encoding]::Unicode.GetString($TIC) -replace '(.+)([A-Z]:[0-9a-zA-Z\\])+','$2'

// https://github.com/austinhyde/wallpaper-go/blob/master/winutils/winutils.go

package wpaper

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

// TODO get the transcodeimagecache byte array from registry
// TODO if possible convert this to c++

func getRegBin(root registry.Key, path string, name string) ([]byte, error) {
	key, err := registry.OpenKey(root, path, registry.READ)
	defer key.Close()
	if err != nil {
		return nil, err
	}
	bin, _, err := key.GetBinaryValue(name)
	return bin, err
}

// github.com/kbinani/win

func getDesktopWallpaper() (string, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.READ)
	defer key.Close()
	if err != nil {
		return "", err
	}
	ret, _, err := key.GetStringValue("Wallpaper")
	if err != nil {
		return "", err
	}
	return ret, nil
}

func getBgColor() (string, error) {
	// HKEY_CURRENT_USER\Control Panel\Colors
	key, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Colors`, registry.READ)
	defer key.Close()
	if err != nil {
		return "", err
	}
	color, _, err := key.GetStringValue("Background")
	return color, err
}

func readTranscodedImageCache(i int) (string, error) {
	name := "TranscodedImageCache"
	withI := name + fmt.Sprintf("_%03d", i)

	if i > 0 {
		return _readTranscodedImageCacheName(withI)
	}

	value, err := _readTranscodedImageCacheName(withI)
	if err == registry.ErrNotExist {
		return _readTranscodedImageCacheName(name)
	}
	return value, err
}

func _readTranscodedImageCacheName(name string) (string, error) {
	bin, err := getRegBin(registry.CURRENT_USER, `Control Panel\Desktop`, name)
	if err == registry.ErrNotExist {
		return "", nil
	} else if err != nil {
		return "", err
	}

	bin = bin[24:]
	n := len(bin) / 2
	data := make([]byte, 0, n)
	for i := 0; i < n && bin[i] != 0; i += 2 {
		data = append(data, bin[i])
	}
	return string(data), nil
}

var (
	setSysColors uintptr
)

// SetSysColors To set colors
// https://github.com/kbinani/win/blob/47348268596a5a7daca78499d4d2e070f7055084/user32.go#L5160
func SetSysColors(cElements int32, lpaElements /*const*/ *int32, lpaRgbValues /*const*/ *colorref) (bool, syscall.Errno) {
	ret1, err := syscall3(setSysColors, 3,
		uintptr(cElements),
		uintptr(unsafe.Pointer(lpaElements)),
		uintptr(unsafe.Pointer(lpaRgbValues)))
	return ret1 != 0, err
}

func syscall3(trap, nargs, a1, a2, a3 uintptr) (uintptr, syscall.Errno) {
	ret, _, err := syscall.Syscall(trap, nargs, a1, a2, a3)
	return ret, err
}

func doLoadLibrary(name string) uintptr {
	lib, _ := syscall.LoadLibrary(name)
	return uintptr(lib)
}

// https://github.com/kbinani/win/blob/47348268596a5a7daca78499d4d2e070f7055084/win_windows.go#L14
func doGetProcAddress(lib uintptr, name string) uintptr {
	addr, _ := syscall.GetProcAddress(syscall.Handle(lib), name)
	return uintptr(addr)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	libuser32 := doLoadLibrary("user32.dll")
	setSysColors = doGetProcAddress(libuser32, "SetSysColors")
	var aElements int32 = ColorBackground
	color := rgb(0, 0, 0)

	d, _ := SetSysColors(1, &aElements, &color)
	log.Println(d)
	// Try first for transcodeimagecache (slideshow)
	data, err := readTranscodedImageCache(0)
	if err != nil {
		log.Fatal(err)
		return
	}
	if data == "" {
		fmt.Println("solid color or wallpaper")
		wallpaper, err := getDesktopWallpaper()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wallpaper)
		color, err := getBgColor()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(color)
	} else {
		fmt.Println(data)
	}
}
