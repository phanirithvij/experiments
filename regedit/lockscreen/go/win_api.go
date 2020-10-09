package main

import "github.com/kbinani/win"

// https://github.com/kbinani/win/blob/47348268596a5a7daca78499d4d2e070f7055084/types.go#L5818
type ushort uint16

type word ushort

// https://github.com/kbinani/win/blob/generator/internal/win/const.go#L273
type dword uint32

// from winuser.h
// #define COLOR_BACKGROUND	1
const (
	ColorBackground int32 = 1
)

// from wingdi.h
// #define RGB(r,g,b)	((COLORREF)(((BYTE)(r)| ((WORD)((BYTE)(g))<<8))| (((DWORD)(BYTE)(b))<<16)))

func rgb(r, g, b int) win.COLORREF {
	rpart := (dword)((byte)(r))
	gpart := (dword)((byte)(g)) << 8
	bpart := (dword)((byte)(b)) << 16
	return (win.COLORREF)((rpart | gpart | bpart))
}
