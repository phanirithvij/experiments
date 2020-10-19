package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
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

func getBgColor() (color.RGBA, error) {
	// HKEY_CURRENT_USER\Control Panel\Colors
	key, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Colors`, registry.READ)
	defer key.Close()
	if err != nil {
		return color.RGBA{}, err
	}
	colorstr, _, err := key.GetStringValue("Background")
	log.Println(colorstr)
	return color.RGBA{}, err
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
		// log.Println("registry not exists")
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

// https://programming.guide/go/define-enumeration-string.html

// WallpaperType Enum for the wallpaper type
type WallpaperType int

const (
	// Picture wallpaper
	Picture WallpaperType = iota
	// SolidColor wallpaper
	SolidColor
	// Slideshow wallpaper
	Slideshow
)

func (w WallpaperType) String() string {
	return [...]string{"Picture", "SolidColor", "Slideshow"}[w]
}

// WallpaperInfo holds the wallpapr information
type WallpaperInfo struct {
	WallpaperType WallpaperType `json:"type"`
	HelpMsg       string        `json:"help"`
	Path          string        `json:"path"`
	SolidColor    color.RGBA    `json:"color"`
	ImageColor    color.RGBA    `json:"imageColor"`
}

func (w WallpaperInfo) String() string {
	argsWithoutProg := os.Args[1:]
	delimiter := ""
	if len(argsWithoutProg) == 0 {
		// delimiter = ""
		// log.Println("Can specify am optional delimiter to format json")
	} else {
		delimiter = argsWithoutProg[0]
	}
	unescapedDelimiter, err := strconv.Unquote(`"` + delimiter + `"`)
	if err != nil {
		log.Println(err)
	}
	// log.Println("unescaped", s, len(s))

	var bytes []byte
	if delimiter != "" {
		bytes, err = json.MarshalIndent(w, "", unescapedDelimiter)
	} else {
		bytes, err = json.Marshal(w)
	}

	if err != nil {
		// https://stackoverflow.com/q/64306027/8608146
		type WallpaperInfoDup WallpaperInfo
		return fmt.Sprintf("%+v\n", WallpaperInfoDup(w))
	}
	bytestr := string(bytes)
	// log.Println("bytestr len", len(bytestr))
	return bytestr
}

func main() {
	log.SetFlags(0)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	winfo := GetWallpaperInfo()
	// https://stackoverflow.com/a/24512194/8608146
	log.Println(winfo)
}

// GetWallpaperInfo returns the Wallpaper info
func GetWallpaperInfo() WallpaperInfo {
	winfo := &WallpaperInfo{}
	wallpaper, err := getDesktopWallpaper()
	if err != nil {
		log.Fatal(err)
	}
	var data string = ""

	if strings.Contains(wallpaper, `AppData\Roaming\Microsoft\Windows\Themes\TranscodedWallpaper`) {
		winfo.WallpaperType = Slideshow
		// Try first for transcodeimagecache (slideshow)
		data, err = readTranscodedImageCache(0)
		if err != nil {
			log.Fatal(err)
		}
		if data == "" {
			// for some reason transcodeimagecache was unreadable
			// print the TranscodedWallpaper path
			winfo.HelpMsg = "Wallpaper's exact path couldn't be determined but this path contains a copy of it"
			winfo.Path = wallpaper
		} else {
			winfo.HelpMsg = "Wallpaper is a slideshow"
			winfo.Path = data
		}

	} else {
		winfo.HelpMsg = "Wallpaper is an image"
		winfo.Path = wallpaper
	}

	if wallpaper == "" && data == "" {
		winfo.HelpMsg = "Not a wallpaper but a solid color"
		winfo.WallpaperType = SolidColor
		color, err := getBgColor()
		if err != nil {
			log.Fatal(err)
		}
		winfo.SolidColor = color
	} else {
		col, err := getImageColor()
		color, err := hexToRGBA(col)
		if err != nil {
			log.Fatal(err)
		}
		winfo.ImageColor = color
	}
	return *winfo
}
