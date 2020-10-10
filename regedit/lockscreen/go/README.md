## TODO

- C:\Windows\Web\Screen
- C:\Users\Rithvij\AppData\Roaming\Microsoft\Windows\Themes\CachedFiles

- [One of the comments here](https://www.deskmodder.de/phpBB3/viewtopic.php?t=16905#p270993) Use google translate extension
- This directory can be read without admin previliges C:\ProgramData\Microsoft\Windows\SystemData\S-1-5-21-1131672954-3644571216-278812857-1001\ReadOnly and it has all the lockscreen wallpapers
- For any directory above this and below and including `...\SystemData` a normal admin access will not work, we need to do so psexec https://superuser.com/a/1416509/1049709

- need to decode `ImagesRootPIDL`
- %USERPROFILE%\AppData\Local\Microsoft\Windows\
- Computer\HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Wallpapers has `BackgroundType` which becomes 0, 1, 2 with Picture, Solid Color, Slideshow respectively

The following places have stuff about the lock screen when in slide show and in picture
- Computer\HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen
<!-- - Computer\HKEY_USERS\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen -->
- HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen

Registry changes

```dart
Regshot 1.9.0 x64 Unicode
Comments: 
Datetime: 2020/10/9 19:39:32  ,  2020/10/9 19:47:15

----------------------------------
Values deleted: 5
----------------------------------
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\ImageId_A: "{A9213B51-4BD3-4D95-9E65-99311032ED79}"
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\OriginalFile_A:  E7 01 14 00 1F 50 E0 4F D0 20 EA 3A 69 10 A2 D8 08 00 2B 30 30 9D 14 00 2E 80 D4 3A AD 24 69 A5 30 45 98 E1 AB 02 F9 41 7A A8 98 00 31 00 00 00 00 00 9E 50 24 3F 11 00 53 43 52 45 45 4E 7E 31 00 00 80 00 09 00 04 00 EF BE 38 4F 25 99 9E 50 24 3F 2E 00 00 00 97 D5 00 00 00 00 05 00 00 00 00 00 00 00 00 00 46 00 00 00 00 00 DF A3 15 01 53 00 63 00 72 00 65 00 65 00 6E 00 73 00 68 00 6F 00 74 00 73 00 00 00 40 00 77 00 69 00 6E 00 64 00 6F 00 77 00 73 00 2E 00 73 00 74 00 6F 00 72 00 61 00 67 00 65 00 2E 00 64 00 6C 00 6C 00 2C 00 2D 00 32 00 31 00 38 00 32 00 33 00 00 00 18 00 25 01 32 00 24 81 21 00 6D 50 0E 78 20 00 53 63 72 65 65 6E 73 68 6F 74 20 28 33 30 29 2E 70 6E 67 00 58 00 09 00 04 00 EF BE 6D 50 0E 78 6D 50 0E 78 2E 00 00 00 FD 64 05 00 00 00 1D 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 66 64 24 00 53 00 63 00 72 00 65 00 65 00 6E 00 73 00 68 00 6F 00 74 
00 20 00 28 00 33 00 30 00 29 00 2E 00 70 00 6E 00 67 00 00 00 22 00 AB 00 00 00 27 00 EF BE 9D 00 00 00 31 53 50 53 B7 9D AE FF 8D 1C FF 43 81 8C 84 40 3A A3 73 2D 81 00 00 00 64 00 00 00 00 1F 00 00 00 37 00 00 00 4D 00 69 00 63 00 72 00 6F 00 73 00 6F 00 66 00 74 00 2E 00 53 00 44 00 4B 00 53 00 61 00 6D 00 70 00 6C 00 65 00 73 00 2E 00 50 00 65 00 72 00 73 00 6F 00 6E 00 61 00 6C 00 69 00 7A 00 61 00 74 00 69 00 6F 00 6E 00 2E 00 43 00 50 00 50 00 5F 00 38 00 77 00 65 00 6B 00 79 00 62 00 33 00 64 00 38 00 62 00 62 00 77 00 65 00 00 00 00 00 00 00 00 00 00 00 00 00 22 00 00 00
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\Details_A: "APPID:Microsoft.SDKSamples.Personalization.CPP_8wekyb3d8bbwe!App"

----------------------------------
Values added: 15
----------------------------------
HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen\SizeX_E: 0x00000780
HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen\SizeY_E: 0x00000438
HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen\CacheFormat_E: 0x00000000
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\ImageId_E: "{7404D834-08A0-4FBF-BBB3-0B119E91EFA4}"
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\OriginalFile_E:  55 01 14 00 1F 50 E0 4F D0 20 EA 3A 69 10 A2 D8 08 00 2B 30 30 9D 3A 00 2E 80 3A CC BF B4 2C DB 4C 42 B0 29 7F E9 9A 87 C6 41 26 00 01 00 26 00 EF BE 11 00 00 00 64 CD 3B BF B4 6E D5 01 1F 97 74 41 71 9E D6 01 19 7E FF C9 73 9E D6 01 14 00 05 01 32 00 35 0E 04 00 3C 51 86 61 20 00 6D 70 76 2D 73 68 6F 74 30 30 30 32 2E 6A 70 67 00 00 52 00 09 00 04 00 EF BE 3C 51 86 61 3D 51 D3 9A 2E 00 00 00 DD 5E 08 00 00 00 31 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 24 46 2F 00 6D 00 70 00 76 00 2D 00 73 00 68 00 6F 00 74 00 30 00 30 00 30 00 32 00 2E 00 6A 00 70 00 67 00 00 00 20 00 93 00 00 00 27 00 EF BE 85 00 00 00 31 53 50 53 B7 9D AE FF 8D 1C FF 43 81 8C 84 40 3A A3 73 2D 69 00 00 00 64 00 00 00 00 1F 00 00 00 2C 00 00 00 77 00 69 00 6E 00 64 00 6F 00 77 00 73 00 2E 00 69 00 6D 00 6D 00 65 00 72 00 73 00 69 00 76 00 65 00 63 00 6F 00 6E 00 74 00 72 00 6F 00 6C 00 70 00 
61 00 6E 00 65 00 6C 00 5F 00 63 00 77 00 35 00 6E 00 31 00 68 00 32 00 74 00 78 00 79 00 65 00 77 00 79 00 00 00 00 00 00 00 00 00 00 00 20 00 00 00
HKU\S-1-5-21-1131672954-3644571216-278812857-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Lock Screen\Details_E: "IMMERSIVE_CPL"

----------------------------------
Values modified: 49
----------------------------------
HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen\: "DZCBYA"
HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\SystemProtectedUserData\S-1-5-21-1131672954-3644571216-278812857-1001\AnyoneRead\LockScreen\: "ECDZBY"
HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Schedule\TaskCache\Tasks\{EC95F45C-0486-40E1-8938-20FE3E377E7D}\DynamicInfo:  03 00 00 00 9B 0E 3C 54 67 6C D6 01 51 18 49 FA 73 9E D6 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Schedule\TaskCache\Tasks\{EC95F45C-0486-40E1-8938-20FE3E377E7D}\DynamicInfo:  03 00 00 00 9B 0E 3C 54 67 6C D6 01 19 54 18 F0 74 9E D6 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
HKLM\SYSTEM\CurrentControlSet\Services\bam\State\UserSettings\S-1-5-21-1131672954-3644571216-278812857-1001\windows.immersivecontrolpanel_cw5n1h2txyewy:  39 5C 93 D6 73 9E D6 01 00 00 00 00 00 00 00 00 01 00 00 00 02 00 00 00
HKLM\SYSTEM\CurrentControlSet\Services\bam\State\UserSettings\S-1-5-21-1131672954-3644571216-278812857-1001\windows.immersivecontrolpanel_cw5n1h2txyewy:  31 F4 C7 FC 74 9E D6 01 00 00 00 00 00 00 00 00 01 00 00 00 02 00 00 00

----------------------------------
Total changes: 74
----------------------------------

```