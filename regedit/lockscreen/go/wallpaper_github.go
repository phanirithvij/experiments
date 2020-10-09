package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kbinani/win"

	"github.com/phanirithvij/experiments/regedit/lockscreen/go/wpaper"

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

func setBgColor(r, g, b int) bool {
	var aElements int32 = wpaper.ColorBackground
	color := wpaper.RGB(r, g, b)
	d := win.SetSysColors(1, &aElements, &color)
	return d
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Try first for transcodeimagecache (slideshow)
	data, err := readTranscodedImageCache(0)
	if err != nil {
		log.Fatal(err)
	}

	wallpaper, err := getDesktopWallpaper()
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(wallpaper, `AppData\Roaming\Microsoft\Windows\Themes\TranscodedWallpaper`) {
		log.Println("wallpaper is a slideshow")
	}
	log.Println(wallpaper)
	if data == "" {
		log.Println("solid color or wallpaper")
		color, err := getBgColor()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(color)
	} else {
		log.Println(data)
	}
}
