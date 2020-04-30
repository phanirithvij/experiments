// this transcodes transcodeimagecache to a string
// look at getwall.ps1
// it's contents are this

// # https://www.winhelponline.com/blog/find-current-wallpaper-file-path-windows-10/
// $TIC=(Get-ItemProperty 'HKCU:\Control Panel\Desktop' TranscodedImageCache -ErrorAction Stop).TranscodedImageCache
// [System.Text.Encoding]::Unicode.GetString($TIC) -replace '(.+)([A-Z]:[0-9a-zA-Z\\])+','$2'

package main

import (
	"fmt"
	"strings"
)

// TODO get the transcodeimagecache byte array from registry
// TODO if possible convert this to python

func main() {
	// this byte array was obtained from

	// $TIC=(Get-ItemProperty 'HKCU:\Control Panel\Desktop' TranscodedImageCache -ErrorAction Stop).TranscodedImageCache
	// $TIC>s.txt

	// then I copy pasted the s.txt contents here
	x := []byte{
		122, 195, 1, 0, 178, 48, 25, 0, 0, 30, 0, 0, 224, 16, 0, 0, 231, 211, 85, 114, 205, 192, 213, 1, 68, 0, 58, 0, 92, 0, 73, 0, 109, 0, 97, 0, 103, 0, 101, 0, 115, 0, 92, 0, 87, 0, 97, 0, 108, 0, 108, 0, 112, 0, 97, 0,
		112, 0, 101, 0, 114, 0, 115, 0, 92, 0, 118, 0, 101, 0, 99, 0, 116, 0, 111, 0, 114, 0, 45, 0, 102, 0, 111, 0, 114, 0, 101, 0, 115, 0, 116, 0, 45, 0, 115, 0, 117, 0, 110, 0, 115, 0, 101, 0, 116, 0, 45, 0, 102, 0, 111,
		0, 114, 0, 101, 0, 115, 0, 116, 0, 45, 0, 115, 0, 117, 0, 110, 0, 115, 0, 101, 0, 116, 0, 45, 0, 102, 0, 111, 0, 114, 0, 101, 0, 115, 0, 116, 0, 45, 0, 119, 0, 97, 0, 108, 0, 108, 0, 112, 0, 97, 0, 112, 0, 101, 0, 114, 0, 45, 0, 98, 0, 51, 0,
		97, 0, 98, 0, 99, 0, 51, 0, 53, 0, 100, 0, 48, 0, 100, 0, 54, 0, 57, 0, 57, 0, 98, 0, 48, 0, 53, 0, 54, 0, 102, 0, 97, 0, 54, 0, 98, 0, 50, 0, 52, 0, 55, 0, 53, 0, 56, 0, 57, 0, 98, 0, 49, 0, 56, 0, 97, 0, 56, 0, 46, 0, 106, 0, 112, 0, 103,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// to remove all 0s from the array
	// they are empty spaces
	// b := x[:] creates a copy
	// b := x[:0] => b is a reference to x
	// https://stackoverflow.com/a/38161077/8608146
	b := x[:0]
	for _, c := range x {
		if c != 0 {
			b = append(b, c)
		}
	}
	// but here len(x) != len(b)
	// eg: if x was [1,0,2,0,3,0,0,0,0] =>
	// b is [1,2,3] and x is [1,2,3,0,3,0,0,0,0]
	r := string(b)
	// the string is of the form eg: jibberishD:\path\to\wallpaper
	// this is all to remove that jibberish
	// split the path by ':'
	split := strings.Split(r, ":")
	drive, path := split[0], split[1]
	// fmt.Println(path)
	// rune is like a special byte which treats unicode characters as a signle byte
	// if we use a byte array then a unicode char can become one or more chars
	runes := []rune(drive)
	// drive i.e. runes has [...jibberish, drive_letter]
	// eg: [122 65533 1 65533 48 25 30 65533 16 65533 65533 85 114 65533 65533 65533 1 68]
	// 68 is 'D'
	// get the last char
	drive = string(runes[len(runes)-1:][0])
	// this is the wallpaper path
	fmt.Println(drive + ":" + path)
}
