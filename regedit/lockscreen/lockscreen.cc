/*
Get a key's value from the registry method #2
cl /EHsc lockscreen.cc keys.cc sid.cc

*/
#pragma comment(lib, "Advapi32.lib")
#include <windows.h>
#include <iostream>
#include <string>
#include "keys.hh"

LONG GetStringRegKey(HKEY hKey, const std::string &strValueName, std::string &strValue, const std::string &strDefaultValue)
{
    strValue = strDefaultValue;
    CHAR szBuffer[512];
    DWORD dwBufferSize = sizeof(szBuffer);
    ULONG nError;
    nError = RegQueryValueExA(hKey, strValueName.c_str(), 0, NULL, (LPBYTE)szBuffer, &dwBufferSize);
    if (ERROR_SUCCESS == nError)
    {
        strValue = szBuffer;
    }
    return nError;
}

string getLockScreenPath()
{
    HKEY hKey;
    LONG lRes = RegOpenKeyExA(
        HKEY_LOCAL_MACHINE,
        getLockScreenRegKey().c_str(),
        // "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Authentication\\LogonUI\\Creative\\S-1-5-21-1131672954-3644571216-278812857-1001\\132327234926016960",
        0,
        KEY_READ,
        &hKey);
    bool bExistsAndSuccess(lRes == ERROR_SUCCESS);
    bool bDoesNotExistsSpecifically(lRes == ERROR_FILE_NOT_FOUND);
    std::string strValueOfBinDir;
    // default-value will be the defalut value
    GetStringRegKey(hKey, "landscapeImage", strValueOfBinDir, "default-value");
    // std::cout << strValueOfBinDir;
    // std::string strKeyDefaultValue;
    // GetStringRegKey(hKey, "", strKeyDefaultValue, "default-value");
    // std::cout << strKeyDefaultValue;

    RegCloseKey(hKey);
    return strValueOfBinDir;
}

int main()
{
    // TODO
    // look at these
    // https://docs.microsoft.com/en-us/uwp/api/windows.system.userprofile.lockscreen.originalimagefile?view=winrt-18362#Windows_System_UserProfile_LockScreen_OriginalImageFile
    // https://docs.microsoft.com/en-us/uwp/api/Windows.System.UserProfile.LockScreen?redirectedfrom=MSDN&view=winrt-18362
    // no endl;
    cout << getLockScreenPath();
    return 0;
}
