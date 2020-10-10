package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"strconv"

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

func getImageColor() (uint64, error) {
	// HKEY_CURRENT_USER\Control Panel\Desktop
	key, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.READ)
	defer key.Close()
	if err != nil {
		return 0, err
	}
	color, _, err := key.GetIntegerValue("ImageColor")
	return color, err
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

// https://github.com/austinhyde/wallpaper-go/blob/master/winutils/winutils.go#L168
// Also forked on my acc https://github.com/phanirithvij/wallpaper-go/blob/master/winutils/winutils.go#L168
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

func hexStringToRGBA(hexstr string) (color.RGBA, error) {
	// var hexuint uint64
	if len(hexstr) == 3 {
		// Odd length hex string
		hexstr = hexstr + hexstr
	}
	bytea, err := hex.DecodeString(hexstr)
	retcol := color.RGBA{}
	if err != nil {
		return retcol, err
	}

	if len(bytea) < 3 {
		return retcol, errors.New("Not a valid hex string")
	}
	if len(bytea) == 4 {
		retcol.A = bytea[0]
		retcol.R = bytea[1]
		retcol.G = bytea[2]
		retcol.B = bytea[3]
	} else {
		// len == 3
		retcol.A = 255
		retcol.R = bytea[0]
		retcol.G = bytea[1]
		retcol.B = bytea[2]
	}
	return retcol, nil
}

// Alternate slow implementation which does int to string and then uses hex decode
func hexToRGBA2(col uint64) (color.RGBA, error) {
	hexstr := strconv.FormatInt((int64)(col), 16)
	colw, err := hexStringToRGBA(hexstr)
	return colw, err
}

func hexToRGBA(col uint64) (color.RGBA, error) {
	// log.Println("0x" + strconv.FormatInt((int64)(col), 16))
	a := (uint8)((col >> 24) & 0xFF) // Extract the AA byte
	r := (uint8)((col >> 16) & 0xFF) // Extract the RR byte
	g := (uint8)((col >> 8) & 0xFF)  // Extract the GG byte
	b := (uint8)((col) & 0xFF)       // Extract the BB byte
	// log.Println(r, g, b, a)
	return color.RGBA{A: a, R: r, G: g, B: b}, nil
}

// func main() {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)

// 	// Try first for transcodeimagecache (slideshow)
// 	data, err := readTranscodedImageCache(0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	wallpaper, err := getDesktopWallpaper()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if strings.Contains(wallpaper, `AppData\Roaming\Microsoft\Windows\Themes\TranscodedWallpaper`) {
// 		log.Println("wallpaper is a slideshow")
// 	}
// 	log.Println(wallpaper)

// 	col, err := getImageColor()
// 	colors, err := hexToRGBA(col)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(colors)

// 	if data == "" {
// 		log.Println("solid color or wallpaper")
// 		color, err := getBgColor()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(color)
// 	} else {
// 		log.Println(data)
// 	}
// }
