#include "lib.h"

void ToggleDesktopIcons()
{
    CComPtr<IFolderView2> spView;
    FindDesktopFolderView(IID_PPV_ARGS(&spView), "IFolderView2");

    DWORD flags = 0;
    ThrowIfFailed(
        spView->GetCurrentFolderFlags(&flags),
        "GetCurrentFolderFlags failed");
    ThrowIfFailed(
        spView->SetCurrentFolderFlags(FWF_NOICONS, flags ^ FWF_NOICONS),
        "SetCurrentFolderFlags failed");
}

// RAII wrapper to initialize/uninitialize COM
struct CComInit
{
    CComInit() { ThrowIfFailed(::CoInitialize(nullptr), "CoInitialize failed"); }
    ~CComInit() { ::CoUninitialize(); }
    CComInit(CComInit const &) = delete;
    CComInit &operator=(CComInit const &) = delete;
};

int main()
{
    try
    {
        CComInit init;

        ToggleDesktopIcons();

        std::cout << "Desktop icons have been toggled.\n";
    }
    catch (std::system_error const &e)
    {
        std::cout << "ERROR: " << e.what() << ", error code: " << e.code() << "\n";
        return 1;
    }

    return 0;
}
