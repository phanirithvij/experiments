'Find current lock screen wallpaper file in Windows 10
'For Windows 10 build 17134 (v1803) and higher
'Created on 14-May-2019 - (c) Ramesh Srinivasan

Option Explicit
Const HKEY_LOCAL_MACHINE = &H80000002
Dim sWallPaper, oReg, strKeyPath, sCurWP
Dim arrSubKeys, subkey, GetOS, GetBuild

GetVersion()
If InStr(LCase(GetOS), "windows 10") = 0 Then WScript.Quit
If CInt(GetBuild) < 17134 Then WScript.Quit
   Dim objFSO: Set objFSO = CreateObject("Scripting.FileSystemObject")
   Dim WshShell : Set WshShell = WScript.CreateObject("WScript.Shell")
   Dim strUser : strUser = CreateObject("WScript.Network").UserName
   Set oReg=GetObject("winmgmts:{impersonationLevel=impersonate}!\\" & _ "." & "\root\default:StdRegProv")
   WScript.Echo oReg
   strKeyPath = "SOFTWARE\Microsoft\Windows\CurrentVersion\Authentication\" & _ "LogonUI\Creative\" + GetSID(strUser)
   oReg.EnumKey HKEY_LOCAL_MACHINE, strKeyPath, arrSubKeys For Each subkey In arrSubKeys sWallPaper = subkey Next strKeyPath = strKeyPath & "\" & sWallPaper oReg.GetStringValue HKEY_LOCAL_MACHINE, strKeyPath, "landscapeImage", sCurWP
   If objFSO.FileExists(sCurWP)
   Then
      Dim sWPTarget sWPTarget = WshShell.ExpandEnvironmentStrings("%userprofile%") & _ "\Pictures\lockscreen_wallpaper.jpg"
      objFSO.CopyFile sCurWP, sWPTarget, True
      WshShell.Run sWPTarget WScript.Sleep 1000
   If MsgBox ("Locate wallpaper file in the Assets folder?", vbYesNo, "Find Wallpaper") = 6
   Then WshShell.run "explorer.exe" & " /select," & sCurWP End
   If Else WScript.Echo("The wallpaper image does not exist on the disk!") WScript.Quit End
   If Function GetSID(UserName)
   Dim DomainName, Result, WMIUser
      If InStr(UserName, "\") > 0
      Then
         DomainName = Mid(UserName, 1, InStr(UserName, "\") - 1)
         UserName = Mid(UserName, InStr(UserName, "\") + 1)
      Else
         DomainName = CreateObject("WScript.Network").UserDomain
      End If
   On Error Resume Next
   Set WMIUser = GetObject("winmgmts:{impersonationlevel=impersonate}!" _
   		& "/root/cimv2:Win32_UserAccount.Domain='" & DomainName & "'" _
   		& ",Name='" & UserName & "'")
   
   If Err.Number = 0 Then
      Result = WMIUser.SID
   Else
      Result = ""
      WScript.Echo "Can't determine the SID. Quitting.."
      WScript.Quit
   End If
   On Error GoTo 0
   GetSID = Result
End

Function GetVersion()
   Dim objWMIService, colOSes, objOS
   Set objWMIService = GetObject("winmgmts:" _
   		& "{impersonationLevel=impersonate}!\\" & "." & "\root\cimv2")
   Set colOSes = objWMIService.ExecQuery("Select * from Win32_OperatingSystem")
   For Each objOS In colOSes
      GetOS =  objOS.Caption
      GetBuild = objOS.BuildNumber
   Next
End Function