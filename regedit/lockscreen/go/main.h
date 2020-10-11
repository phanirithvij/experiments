#ifndef _MAIN___
#define _MAIN___
#pragma comment(lib, "Ole32.lib")
#include <windows.h>
#include <shobjidl.h>
#include <cstdio>
#include <cstdlib>

#ifndef DLLAPI
#define DLLAPI extern "C" __declspec(dllexport)
#endif

typedef struct
{
  char *SlideshowPath;
  char *WallpaperPath;
} WallpapersInfo;

DLLAPI bool system_wallpaper_info(WallpapersInfo *&info);
DLLAPI void lwpstrToString(LPWSTR lwpstr, char *&buffer);
// DLLAPI WallpapersInfo *system_wallpaper_info(WallpapersInfo *info);
#endif