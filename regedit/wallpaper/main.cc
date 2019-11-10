#include "main.h"
#include <iostream>

// cl wallpaper.cc main.cc /EHsc

int main()
{
    // https://stackoverflow.com/a/9840028/8608146
    std::wcout << system_wallpaper();
    return 0;
}