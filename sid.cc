/*
This gets the current process' user's sid
// https://stackoverflow.com/q/2686096/8608146
// https://www.codeproject.com/Articles/14828/How-To-Get-Process-Owner-ID-and-Current-User-SID   comments has the cpp version
// I web archived it https://web.archive.org/web/20160502161610/http://www.codeproject.com/Articles/14828/How-To-Get-Process-Owner-ID-and-Current-User-SID?msg=4038457#xx4038457xx
cl /EHsc sid.cc /link Advapi32.lib
*/
#include <iostream>
#include <string>
#include <windows.h>
#include <sddl.h>
using namespace std;

PSID GetUserSIDFromPID(DWORD dwProcessID)
{
    HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION, FALSE, dwProcessID);
    if (hProcess)
    {
        HANDLE token = nullptr;
        if (!OpenProcessToken(hProcess, TOKEN_QUERY, &token))
        {
            printf("OpenProcessToken failed. GetLastError=%d", GetLastError());
        }
        else
        {
            DWORD iTokenInfLength = 0;
            if (!GetTokenInformation(token, TokenUser, 0, iTokenInfLength, &iTokenInfLength))
            {
                if (GetLastError() != ERROR_INSUFFICIENT_BUFFER)
                {
                    printf("GetTokenInformation failed. GetLastError=%d", GetLastError());
                }
                else
                {
                    PVOID pTokenInfo = new char[iTokenInfLength];
                    if (!GetTokenInformation(token, TokenUser, pTokenInfo, iTokenInfLength, &iTokenInfLength))
                    {
                        printf("GetTokenInformation failed. GetLastError=%d", GetLastError());
                    }
                    else
                    {
                        TOKEN_USER *pTokenUser = (TOKEN_USER *)pTokenInfo;
                        return pTokenUser->User.Sid;
                    }
                    delete pTokenInfo;
                }
            }
        }
        CloseHandle(hProcess);
    }
    return nullptr;
}

string GetProcessSID(DWORD dwProcessID)
{
    LPSTR pstrSid = nullptr;
    ConvertSidToStringSidA(GetUserSIDFromPID(dwProcessID), &pstrSid);

    string wstrSID;
    if (pstrSid != nullptr)
    {
        wstrSID = pstrSid;
    }

    LocalFree(pstrSid);

    return wstrSID;
}
int main(int argc, char *argv[])
{
    // PID of current process
    DWORD dwPID = GetCurrentProcessId();
    std::cout << GetProcessSID(dwPID) << '\n';
    return 0;
}