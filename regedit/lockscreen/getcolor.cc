#include <windows.h>
#include <stdio.h>
#pragma comment(lib, "user32.lib")

void main()
{
    int aElements[1] = {COLOR_BACKGROUND};
    DWORD aOldColors[1];
    DWORD aNewColors[1];

    // Get the current color of the window background.

    aOldColors[0] = GetSysColor(aElements[0]);

    printf("Current bg color: {0x%x, 0x%x, 0x%x}\n",
           GetRValue(aOldColors[0]),
           GetGValue(aOldColors[0]),
           GetBValue(aOldColors[0]));

    // Define new colors for the elements

    aNewColors[0] = RGB(0x80, 0x00, 0x80); // dark purple

    printf("\nNew bg color: {0x%x, 0x%x, 0x%x}\n",
           GetRValue(aNewColors[0]),
           GetGValue(aNewColors[0]),
           GetBValue(aNewColors[0]));

    // Set the elements defined in aElements to the colors defined
    // in aNewColors

    SetSysColors(1, aElements, aNewColors);

    printf("\nWindow background havs been changed.\n");
    // printf("Reverting to previous colors in 10 seconds...\n");

    // Sleep(10000);

    // Restore the elements to their original colors

    // SetSysColors(1, aElements, aOldColors);
}
