// https://github.com/aont/spotlight_wallpaper/blob/c44f801630de2736677af096d613d3cfb12e3cde/system_wallpaper.cpp
#pragma comment(lib, "Ole32.lib")
#include <windows.h>
#include <shobjidl.h>
#include <cstdio>
#include <cstdlib>
#include <iostream>
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

DLLAPI LPWSTR system_wallpaper(WallpapersInfo *& info)
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

    // https://github.com/psychoria/MFCClasses/blob/5698c849f4206e535f068d72fdf16c00e5544998/PGWallpaper.cpp#L125
    IShellItemArray *pIShell = nullptr;
    hr = pDesktopWallpaper->GetSlideshow(&pIShell);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetSlideShow returned 0x%08x", hr);
        return NULL;
    }

    DWORD dwd = 0;
    hr = pIShell->GetCount(&dwd);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] pIShell->GetCount returned 0x%08x", hr);
        return NULL;
    }

    IShellItem *pIShellItem = nullptr;
    hr = pIShell->GetItemAt(0, &pIShellItem);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] pIShell->GetItemAt returned 0x%08x", hr);
        return NULL;
    }

    // https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-sigdn
    LPWSTR slideshow_path;
    pIShellItem->GetDisplayName(SIGDN_FILESYSPATH, &slideshow_path);

    info->SlideshowPath = slideshow_path;
    info->WallpaperPath = wallpaper_wcs;

    return NULL;
}
