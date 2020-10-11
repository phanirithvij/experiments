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

DLLAPI bool system_wallpaper_info(WallpapersInfo *&info)
{

    HRESULT hr = CoInitialize(NULL);
    if (FAILED(hr))
    {
        fprintf(stderr, "[error] CoInitialize returned 0x%08x\n", hr);
        return false;
    }
    CoUninitializeOnExit cuoe;

    IDesktopWallpaper *pDesktopWallpaper = NULL;
    hr = CoCreateInstance(__uuidof(DesktopWallpaper), NULL, CLSCTX_ALL, IID_PPV_ARGS(&pDesktopWallpaper));
    if (FAILED(hr))
    {
        fprintf(stderr, "[error] CoCreateInstance(__uuidof(DesktopWallpaper)) returned 0x%08x\n", hr);
        return false;
    }
    ReleaseOnExit releaseDesktopWallpaper((IUnknown *)pDesktopWallpaper);

    LPWSTR monitorID;
    hr = pDesktopWallpaper->GetMonitorDevicePathAt(0, &monitorID);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetMonitorDevicePathAt returned 0x%08x\n", hr);
        return false;
    }

    LPWSTR wallpaper_wcs;
    hr = pDesktopWallpaper->GetWallpaper(monitorID, &wallpaper_wcs);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetWallpaper returned 0x%08x\n", hr);
        return false;
    }
    lwpstrToString(wallpaper_wcs, info->WallpaperPath);

    // https://social.msdn.microsoft.com/Forums/en-US/edc2e1de-c7c6-4bef-becb-cf4924165551/decode-encrypted-path-from-slideshowdirectorypath1?forum=windowsgeneraldevelopmentissues
    // The problem is,  IShellItemArray items = wallpaper.GetSlideshow(); where IDesktopWallpaper wallpaper = (IDesktopWallpaper)(new DesktopWallpaperClass()); gives me only the correct path, when slideshow is turned on. But for activating it (wallpaper.SetSlideshow(items)) , i need stored path from SlideshowDirectoryPath1.
    // https://github.com/psychoria/MFCClasses/blob/5698c849f4206e535f068d72fdf16c00e5544998/PGWallpaper.cpp#L125
    IShellItemArray *pIShell = nullptr;
    hr = pDesktopWallpaper->GetSlideshow(&pIShell);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] IDesktopWallpaper::GetSlideShow returned 0x%08x\n", hr);
        return false;
    }

    DWORD dwd = 0;
    hr = pIShell->GetCount(&dwd);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] pIShell->GetCount returned 0x%08x\n", hr);
        return false;
    }

    IShellItem *pIShellItem = nullptr;
    hr = pIShell->GetItemAt(0, &pIShellItem);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] pIShell->GetItemAt returned 0x%08x\n", hr);
        return false;
    }

    // https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-sigdn
    LPWSTR slideshow_path;
    hr = pIShellItem->GetDisplayName(SIGDN_FILESYSPATH, &slideshow_path);
    if (hr != S_OK)
    {
        fprintf(stderr, "[error] pIShell->GetDisplayName returned 0x%08x\n", hr);
        return false;
    }
    lwpstrToString(slideshow_path, info->SlideshowPath);

    // TODO free stuff properly
    CoTaskMemFree(wallpaper_wcs);
    CoTaskMemFree(slideshow_path);
    slideshow_path = nullptr;
    wallpaper_wcs = nullptr;

    return true;
}

void lwpstrToString(LPWSTR lwpstr, char *&buffer)
{
    int size = wcslen(lwpstr);
    buffer = new char[size];
    // First arg is the pointer to destination char, second arg is
    // the pointer to source wchar_t, last arg is the size of char buffer
    size = wcstombs(buffer, lwpstr, MAX_PATH);
    // TODO this line not working so memory leaks
    // delete[] buffer;
}
