// https://github.com/aont/spotlight_wallpaper/blob/c44f801630de2736677af096d613d3cfb12e3cde/system_wallpaper.cpp
#pragma comment(lib, "Ole32.lib")
#include <windows.h>
#include <shobjidl.h>
#include <cstdio>
#include <cstdlib>
#include "main.h"

// ref: https://matthewvaneerde.wordpress.com/2012/10/10/changing-the-desktop-wallpaper-using-idesktopwallpaper/
//  https://github.com/mvaneerde/blog/blob/develop/desktopwallpaper/desktopwallpaper/main.cpp

class CoUninitializeOnExit
{
public:
    CoUninitializeOnExit() {}
    ~CoUninitializeOnExit() { CoUninitialize(); }
};

class ReleaseOnExit
{
public:
    ReleaseOnExit(IUnknown *p) : m_p(p) {}
    ~ReleaseOnExit()
    {
        if (NULL != m_p)
        {
            m_p->Release();
        }
    }

private:
    IUnknown *m_p;
};

DLLAPI LPWSTR system_wallpaper()
{

    HRESULT hr = CoInitialize(NULL);
    if (FAILED(hr))
    {
        fprintf(stderr, "[error] CoInitialize returned 0x%08x", hr);
        return NULL;
    }
    CoUninitializeOnExit cuoe;

    IDesktopWallpaper *pDesktopWallpaper = NULL;
    hr = CoCreateInstance(__uuidof(DesktopWallpaper), NULL, CLSCTX_ALL, IID_PPV_ARGS(&pDesktopWallpaper));
    if (FAILED(hr))
    {
        fprintf(stderr, "[error] CoCreateInstance(__uuidof(DesktopWallpaper)) returned 0x%08x", hr);
        return NULL;
    }
    ReleaseOnExit releaseDesktopWallpaper((IUnknown *)pDesktopWallpaper);

    LPWSTR monitorID;
    hr = pDesktopWallpaper->GetMonitorDevicePathAt(0, &monitorID);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetMonitorDevicePathAt returned 0x%08x", hr);
        return NULL;
    }

    LPWSTR wallpaper_wcs;
    hr = pDesktopWallpaper->GetWallpaper(monitorID, &wallpaper_wcs);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetWallpaper returned 0x%08x", hr);
        return NULL;
    }

    return wallpaper_wcs;

    // int const wcs_len = wcslen(wallpaper_wcs);
    // // int mbs_len = wcs_len*sizeof(WCHAR);
    // int const mbs_len = WideCharToMultiByte(0, 0, wallpaper_wcs, wcs_len, NULL, 0, NULL, NULL);
    // if ( mbs_len==0 ) {
    //     fprintf(stderr, "[error] WideCharToMultiByte failed\n");
    //     return NULL;
    // }
    // char* wallpaper_mbs = (char*)malloc(mbs_len);
    // int const ret = WideCharToMultiByte(0, 0, wallpaper_wcs, wcs_len, wallpaper_mbs, mbs_len, NULL, NULL);
    // if ( ret==0 ) {
    //     fprintf(stderr, "[error] WideCharToMultiByte failed\n");
    //     return NULL;
    // }

    // // fwrite(wallpaper_mbs, mbs_len, 1, stdout);
    // printf("%s\n", wallpaper_mbs);
    // free(wallpaper_mbs);

    // fprintf(stderr, "acp: %u\n", GetACP());
    // return 0;
}