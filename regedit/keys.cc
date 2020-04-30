/*
This code recursively prints all keys, subkeys under a given key in the Windows Registry

// vcvars32.bat Didn't work
vcvars64.bat
cl /EHsc keys.cc /link Advapi32.lib
*/

#pragma comment(lib, "Advapi32.lib")
#include <windows.h>
#include <stdio.h>

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

void EnumerateSubKeys(HKEY RootKey, char *subKey, unsigned int tabs = 0)
{
    HKEY hKey;
    DWORD cSubKeys;     //Used to store the number of Subkeys
    DWORD maxSubkeyLen; //Longest Subkey name length
    DWORD cValues;      //Used to store the number of Subkeys
    DWORD maxValueLen;  //Longest Subkey name length
    DWORD retCode;      //Return values of calls

    RegOpenKeyExA(RootKey, subKey, 0, KEY_ALL_ACCESS, &hKey);

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
                for (int i = 0; i < tabs; i++)
                    printf("\t");
                printf("(%d) %s\n", i + 1, currentSubkey);

                char *subKeyPath = new char[currentSubLen + strlen(subKey)];
                sprintf(subKeyPath, "%s\\%s", subKey, currentSubkey);
                EnumerateSubKeys(RootKey, subKeyPath, (tabs + 1));
            }
        }
    }
    else
    {
        EnumerateValues(hKey, cValues);
    }

    RegCloseKey(hKey);
}

int main()
{
    EnumerateSubKeys(HKEY_LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Authentication\\LogonUI\\Creative\\S-1-5-21-1131672954-3644571216-278812857-1001");
    return 0;
}
