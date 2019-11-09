Set Shell = CreateObject("WScript.Shell")

' change to TranscodedImageCache_001 for second monitor and so on
openWallpaper("HKCU\Control Panel\Desktop\TranscodedImageCache_000")

function openWallpaper(regKey) 
  arr = Shell.RegRead(regKey)
  a=arr
  fullPath = ""
  consequtiveZeroes = 0

  For I = 24 To Ubound(arr)
    if consequtiveZeroes > 1 then
      exit for
    end if

    a(I) = Cint(arr(I))

    if a(I) > 1 then
      fullPath = fullPath & Chr(a(I))
      consequtiveZeroes = 0
    else
      consequtiveZeroes = consequtiveZeroes + 1
    end if
  Next

  Shell.run fullPath
end function

Wscript.Quit