#include <windows.h>
#include <stdio.h>
#include <string>

using namespace std;

void EnumerateValues(HKEY hKey, DWORD numValues);
string EnumerateSubKeys(HKEY RootKey, string subKey, unsigned int tabs);
string getLockScreenRegKey();
