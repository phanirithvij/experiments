/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: wallpaper.i

package wallpaper

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef int intgo;
typedef unsigned int uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
extern void _wrap_Swig_free_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_wallpaper_1a7abfaf8ddc9694(swig_intgo arg1);
extern void _wrap_WallpapersInfo_SlideshowPath_set_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1, swig_type_1 arg2);
extern swig_type_2 _wrap_WallpapersInfo_SlideshowPath_get_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1);
extern void _wrap_WallpapersInfo_WallpaperPath_set_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1, swig_type_3 arg2);
extern swig_type_4 _wrap_WallpapersInfo_WallpaperPath_get_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1);
extern uintptr_t _wrap_new_WallpapersInfo_wallpaper_1a7abfaf8ddc9694(void);
extern void _wrap_delete_WallpapersInfo_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1);
extern _Bool _wrap_system_wallpaper_info_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1);
extern void _wrap_lwpstrToString_wallpaper_1a7abfaf8ddc9694(uintptr_t arg1, swig_type_5 arg2);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_wallpaper_1a7abfaf8ddc9694(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrWallpapersInfo uintptr

func (p SwigcptrWallpapersInfo) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrWallpapersInfo) SwigIsWallpapersInfo() {
}

func (arg1 SwigcptrWallpapersInfo) SetSlideshowPath(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_WallpapersInfo_SlideshowPath_set_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0), *(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrWallpapersInfo) GetSlideshowPath() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_WallpapersInfo_SlideshowPath_get_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrWallpapersInfo) SetWallpaperPath(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_WallpapersInfo_WallpaperPath_set_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrWallpapersInfo) GetWallpaperPath() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_WallpapersInfo_WallpaperPath_get_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func NewWallpapersInfo() (_swig_ret WallpapersInfo) {
	var swig_r WallpapersInfo
	swig_r = (WallpapersInfo)(SwigcptrWallpapersInfo(C._wrap_new_WallpapersInfo_wallpaper_1a7abfaf8ddc9694()))
	return swig_r
}

func DeleteWallpapersInfo(arg1 WallpapersInfo) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_WallpapersInfo_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0))
}

type WallpapersInfo interface {
	Swigcptr() uintptr
	SwigIsWallpapersInfo()
	SetSlideshowPath(arg2 string)
	GetSlideshowPath() (_swig_ret string)
	SetWallpaperPath(arg2 string)
	GetWallpaperPath() (_swig_ret string)
}

func System_wallpaper_info(arg1 WallpapersInfo) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (bool)(C._wrap_system_wallpaper_info_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func LwpstrToString(arg1 LPWSTR, arg2 string) {
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	C._wrap_lwpstrToString_wallpaper_1a7abfaf8ddc9694(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}


type SwigcptrLPWSTR uintptr
type LPWSTR interface {
	Swigcptr() uintptr;
}
func (p SwigcptrLPWSTR) Swigcptr() uintptr {
	return uintptr(p)
}

