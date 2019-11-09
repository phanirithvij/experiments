' Copyright © 2013 Dwight Grant.  All rights reserved.
' Filename:  Win8_WP_Curr_Image_Name-Folder.txt
' Version:  2.00.01
'
' Purpose:  If (in Windows 8) running Desktop Wallpaper Slideshow:
'              To Display the Name of the Current Image (opt=1)
'                                      or
'              Display the Folder with the Current Image Selected (opt=2).
'
' This script reads and decodes registry key:
'      HKEY_CURRENT_USER\Control Panel\Desktop\TranscodedImageCache
'          and displays results on desktop
'
' How to Use  -  Creation:
' 1. Save Text file at a location of your choosing.
' 2. Open Text file with "Notepad" and "Save As" same name with ".vbs" file type.
'       You will now have both  "Win8_WP_Curr_Image_Name-Folder.txt" and
'                               "Win8_WP_Curr_Image_Name-Folder.vbs".
' 3. Create desktop shortcut to "Win8_WP_Curr_Image_Name-Folder.vbs".
'
' Operation:  Double Click on the Desktop Icon Created,
'               Executes "Microsoft Windows Based Script Host" and
'                 will display full path name of Wallpaper file.
'
'***  Author: Dwight Grant  **** Revised: Nov. 27, 2013 ***
' based upon idea from Ramesh Srinivasan in program "WPTargetDir.vbs" for Win 7
' & revisions suggested by FleetCommand.
'
' ***  Please note:  It is not unicode compliant - Path name needs to be ASCII to display properly.
'        If anyone has sugestions as to how to make it compliant, please explain, and I will try to 
'             incorporate it into the next version.
'  **********************************************************
Set Shell    = CreateObject("WScript.Shell")

strEr1       = "Error "
strSingle    = " "
strSelect    = " /select,"
strExplor    = "   "
opt          =  2         '1= Display File Name Only  -  2=Display Folder w/ File Name Selected

strPath      = ""                          'Path Name w/ leading blanks removed
sQ1          = """"                        'A QUOTE mark
Results      = "   "

On Error Resume Next
arr          = Shell.RegRead("HKCU\Control Panel\Desktop\TranscodedImageCache")
If Err.Number <> 0 Then
     strEr1  =  strEr1 & CStr(Err.Number)     'Set error display string
     msgbox strEr1,,"Win8 WP Curr Image Name"  'display error
     WScript.Quit
End If
On Error Goto 0

a            = arr
For I = LBound(arr) To Ubound(arr)         'Pull data from "arr" and convert to integer 
a(I) = Cint(arr(I))                        'Store integer in array "a"
   if I > 23 then                          'Disregard the first 23 characters
       strSingle = Chr(a(I))               'Move byte in array "a" to "strSingle”
       if a(I) > 0 then                    'If byte > zero, use it, else ignore
           strPath = strPath & strSingle   'Add character to string for display
       end if
  end if
Next
        
'  **********************************************************
if opt = 1 then
     msgbox strPath,,"The Wallpaper File Name is"   'Display results on desktop screen
end if
if opt = 2 then
     Results = sQ1 & strPath & sQ1
     strExplor = strSelect & Results 
          'msgbox strExplor,,"The String Passed to Explorer is"     'Diagnostic Display
     return = Shell.run("explorer.exe" & strExplor,,true)
end if
'  **********************************************************
Wscript.Quit
