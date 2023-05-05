#pragma once
#pragma comment(lib, "Shell32.lib")
#pragma comment(lib, "Ole32.lib")

#include <ShlObj.h>    // Shell API
#include <atlcomcli.h> // CComPtr & Co.
#include <string>
#include <iostream>
#include <system_error>

void FindDesktopFolderView(REFIID riid, void **ppv, std::string const &interfaceName);

template <typename T>
void ThrowIfFailed(HRESULT hr, T &&msg);
