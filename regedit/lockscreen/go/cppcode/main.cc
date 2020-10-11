#include <iostream>
#include "shlobj.h"
#include "shlobj_core.h"
#include <Objbase.h>
#include "main.h"

// #pragma comment(lib, "shell32.lib")
// #pragma comment(lib, "Ole32.lib")

/* 
    vcvars64.bat
    cl wallpaper.cc main.cc /EHsc
    .\wallpaper.exe
*/
using namespace std;

int main()
{
    // https://stackoverflow.com/a/9840028/8608146
    WallpapersInfo info;
    WallpapersInfo *infoptr = &info;
    bool ret = system_wallpaper_info(infoptr);

    std::cout << info.WallpaperPath << std::endl;
    if (ret)
    {
        // If sideshow not active ret WILL be FALSE

        // will crash if ret is false and we access this
        std::cout << info.SlideshowPath << std::endl;

        // These free won't free
        // free(infoptr->SlideshowPath);
        // free(infoptr->WallpaperPath);

        // these free will crash the program
        // free(infoptr);

        infoptr->SlideshowPath = nullptr;
        infoptr->WallpaperPath = nullptr;
        infoptr = nullptr;
    }

    return 0;
}

bool GetDocsPath(char *pszDesktopPath);
LPITEMIDLIST getPIDLFromPath(string pszPath)
{
    return ILCreateFromPathA(pszPath.c_str());
}

void examplePDIL()
{
    char buf[MAX_PATH];
    bool ret = GetDocsPath(buf);
    if (!ret)
    {
        cerr << "Failed to get docs path" << endl;
    }
    string docs = string(buf);
    cout << docs << endl;
    LPITEMIDLIST pdil = getPIDLFromPath(docs);
    cout << pdil->mkid.abID << endl;
    cout << pdil->mkid.cb << endl;
}

// modified up on
// https://github.com/Peakchen/xWinManager/blob/49f291dad71c02e948c40d94078f7a485d4250d7/src/quicksettings.cpp#L15
bool GetDocsPath(char *pszDesktopPath)
{
    LPITEMIDLIST ppidl = NULL;
    if (SHGetSpecialFolderLocation(NULL, CSIDL_MYDOCUMENTS, &ppidl) == S_OK)
    {
        BOOL flag = SHGetPathFromIDListA(ppidl, pszDesktopPath);
        CoTaskMemFree(ppidl);
        return bool(flag);
    }
    return false;
}