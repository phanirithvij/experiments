#include "lib.h"

// Throw a std::system_error if the HRESULT indicates failure.
template <typename T>
void ThrowIfFailed(HRESULT hr, T &&msg)
{
    if (FAILED(hr))
        throw std::system_error{hr, std::system_category(), std::forward<T>(msg)};
}

// Query an interface from the desktop shell view.
void FindDesktopFolderView(REFIID riid, void **ppv, std::string const &interfaceName)
{
    CComPtr<IShellWindows> spShellWindows;
    ThrowIfFailed(
        spShellWindows.CoCreateInstance(CLSID_ShellWindows),
        "Failed to create IShellWindows instance");

    CComVariant vtLoc(CSIDL_DESKTOP);
    CComVariant vtEmpty;
    long lhwnd;
    CComPtr<IDispatch> spdisp;
    ThrowIfFailed(
        spShellWindows->FindWindowSW(
            &vtLoc, &vtEmpty, SWC_DESKTOP, &lhwnd, SWFO_NEEDDISPATCH, &spdisp),
        "Failed to find desktop window");

    CComQIPtr<IServiceProvider> spProv(spdisp);
    if (!spProv)
        ThrowIfFailed(E_NOINTERFACE, "Failed to get IServiceProvider interface for desktop");

    CComPtr<IShellBrowser> spBrowser;
    ThrowIfFailed(
        spProv->QueryService(SID_STopLevelBrowser, IID_PPV_ARGS(&spBrowser)),
        "Failed to get IShellBrowser for desktop");

    CComPtr<IShellView> spView;
    ThrowIfFailed(
        spBrowser->QueryActiveShellView(&spView),
        "Failed to query IShellView for desktop");

    ThrowIfFailed(
        spView->QueryInterface(riid, ppv),
        "Could not query desktop IShellView for interface " + interfaceName);
}
