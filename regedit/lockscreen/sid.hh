#include <iostream>
#include <string>
#include <windows.h>
#include <sddl.h>

using namespace std;

PSID GetUserSIDFromPID(DWORD dwProcessID);
string GetProcessSID(DWORD dwProcessID);
