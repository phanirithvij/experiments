/*
This code recursively prints all keys, subkeys under a given key in the Windows Registry

// vcvars32.bat Didn't work
vcvars64.bat
cl /EHsc keys.cc sid.cc /link Advapi32.lib
*/

// with this line /link Advapi32.lib is obsolete
#pragma comment(lib, "Advapi32.lib")
#include "sid.hh"
#include "keys.hh"

using namespace std;

void EnumerateValues(HKEY hKey, DWORD numValues)
{
    DWORD dwIndex = 0;
    LPSTR valueName = new CHAR[64];
    DWORD valNameLen = 64;
    DWORD numback = numValues;

    for (int i = 0; i < numValues; i++)
    {
        // https://stackoverflow.com/a/50766807/8608146
        // Last 4 must be nullptr
        RegEnumValueA(hKey,
                      dwIndex,
                      valueName,
                      &valNameLen,
                      nullptr,
                      nullptr,
                      nullptr,
                      nullptr);
        dwIndex++;

        if (i > numback)
        {
            RegCloseKey(hKey);
            printf("Inf loop exiting...\n");
            exit(-1);
        }
    }
    RegCloseKey(hKey);
}

string EnumerateSubKeys(HKEY RootKey, string subKey, unsigned int tabs = 0)
{
    string lastSubKey;

    HKEY hKey;
    DWORD cSubKeys;     //Used to store the number of Subkeys
    DWORD maxSubkeyLen; //Longest Subkey name length
    DWORD cValues;      //Used to store the number of Subkeys
    DWORD maxValueLen;  //Longest Subkey name length
    DWORD retCode;      //Return values of calls

    RegOpenKeyExA(RootKey, subKey.c_str(), 0, KEY_ALL_ACCESS, &hKey);

    RegQueryInfoKey(hKey,          // key handle
                    nullptr,       // buffer for class name
                    nullptr,       // size of class string
                    nullptr,       // reserved
                    &cSubKeys,     // number of subkeys
                    &maxSubkeyLen, // longest subkey length
                    nullptr,       // longest class string
                    &cValues,      // number of values for this key
                    &maxValueLen,  // longest value name
                    nullptr,       // longest value data
                    nullptr,       // security descriptor
                    nullptr);      // last write time

    if (cSubKeys > 0)
    {
        char currentSubkey[MAX_PATH];

        for (int i = 0; i < cSubKeys; i++)
        {
            DWORD currentSubLen = MAX_PATH;

            retCode = RegEnumKeyExA(hKey,           // Handle to an open/predefined key
                                    i,              // Index of the subkey to retrieve.
                                    currentSubkey,  // buffer to receives the name of the subkey
                                    &currentSubLen, // size of that buffer
                                    nullptr,        // Reserved
                                    nullptr,        // buffer for class string
                                    nullptr,        // size of that buffer
                                    nullptr);       // last write time

            if (retCode == ERROR_SUCCESS)
            {
                // for (int i = 0; i < tabs; i++)
                //     printf("\t");
                // printf("(%d) %s\n", i + 1, currentSubkey);

                char *subKeyPath = new char[currentSubLen + subKey.size()];
                // sprintf(subKeyPath, "%s\\%s", subKey.c_str(), currentSubkey);
                lastSubKey = string(currentSubkey);
                // recursion for going down
                // not needed for my use
                // EnumerateSubKeys(RootKey, subKeyPath, (tabs + 1));
            }
        }
    }
    else
    {
        EnumerateValues(hKey, cValues);
    }

    RegCloseKey(hKey);

    // remove this if using recusion
    return lastSubKey;
}

string getLockScreenRegKey()
{
    // https://www.winhelponline.com/blog/find-file-name-lock-screen-image-current-displayed/

    // In Windows 10 versions 1803 and higher, the current lock screen wallpaper image is stored in string values (REG_SZ) namely landscapeImage and portraitImage, under the following registry key:
    // HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Authentication\LogonUI\Creative\<your SID>\<random-key-name>

    // TODO
    /* In Windows 10 version 1709 and earlier, the lock screen image (Windows Spotlight) file name for the currently displayed landscape and portrait assets are stored in the following registry key:
        HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\Creative
    */

    // get sid
    // PID of current process
    DWORD dwPID = GetCurrentProcessId();
    string sid = GetProcessSID(dwPID);
    // cout << sid << '\n';
    // get the last subKey out of availabe 3 (in my pc)
    string lastSubKey = EnumerateSubKeys(
        HKEY_LOCAL_MACHINE,
        "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Authentication\\LogonUI\\Creative\\" + sid);

    return "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Authentication\\LogonUI\\Creative\\" + sid + "\\" + lastSubKey;
}
