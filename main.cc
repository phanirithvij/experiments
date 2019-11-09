#include <iostream>
#include <string>
#include <exception>
#include <windows.h>
/*
    Get a key's value from registry
    // https://stackoverflow.com/a/22954182/8608146
    cl /EHsc main.cc /link Advapi32.lib
*/

/*! \brief                          Returns a value from HKLM as string.
    \exception  std::runtime_error  Replace with your error handling.
*/
std::wstring GetStringValueFromHKLM(const std::wstring &regSubKey, const std::wstring &regValue)
{
    size_t bufferSize = 0xFFF; // If too small, will be resized down below.
    std::wstring valueBuf;     // Contiguous buffer since C++11.
    valueBuf.resize(bufferSize);
    DWORD cbData = static_cast<DWORD>(bufferSize);
    LONG rc = RegGetValueW(
        HKEY_LOCAL_MACHINE,
        regSubKey.c_str(),
        regValue.c_str(),
        RRF_RT_REG_SZ,
        nullptr,
        static_cast<void *>(&valueBuf.at(0)),
        &cbData);
    while (rc == ERROR_MORE_DATA)
    {
        // Get a buffer that is big enough.
        cbData /= sizeof(wchar_t);
        if (cbData > static_cast<DWORD>(bufferSize))
        {
            bufferSize = static_cast<size_t>(cbData);
        }
        else
        {
            bufferSize *= 2;
            cbData = static_cast<DWORD>(bufferSize);
        }
        valueBuf.resize(bufferSize);
        rc = RegGetValueW(
            HKEY_LOCAL_MACHINE,
            regSubKey.c_str(),
            regValue.c_str(),
            RRF_RT_REG_SZ,
            nullptr,
            static_cast<void *>(&valueBuf.at(0)),
            &cbData);
    }
    if (rc == ERROR_SUCCESS)
    {
        valueBuf.resize(static_cast<size_t>(cbData / sizeof(wchar_t)));
        return valueBuf;
    }
    else
    {
        std::cout << rc;
        throw std::runtime_error("Windows system error code: " + std::to_string(rc));
    }
}

int main()
{
    std::wstring regSubKey;
#ifdef _WIN64 // Manually switching between 32bit/64bit for the example. Use dwFlags instead.
    regSubKey = L"SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Authentication\\LogonUI\\Creative\\S-1-5-21-1131672954-3644571216-278812857-1001\\132177734358952544";
#else
    regSubKey = L"SOFTWARE\\Company Name\\Application Name\\";
#endif
    std::wstring regValue(L"landscapeImage");
    std::wstring valueFromRegistry;
    try
    {
        valueFromRegistry = GetStringValueFromHKLM(regSubKey, regValue);
    }
    catch (std::exception &e)
    {
        std::cerr << e.what();
    }
    std::wcout << valueFromRegistry;
}