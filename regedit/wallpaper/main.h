#ifndef _MAIN___
#define _MAIN___
#pragma comment( lib, "Ole32.lib" )
#include <windows.h>
#include <shobjidl.h>
#include <cstdio>
#include <cstdlib>

#ifndef DLLAPI
  #define DLLAPI extern "C" __declspec(dllexport)
#endif

DLLAPI LPWSTR system_wallpaper();
#endif
