#include <WinAPIShellEx.au3>

$sTestPath = "C:\Windows" ; Must be an existing folder
ConsoleWrite( @CRLF & "Test path: $sTestPath = " & $sTestPath & @CRLF & @CRLF )

;$sEncodedPidl = _EncryptedPIDLEncode2( $sTestPath )
$sEncodedPidl = "2FAFA8BUg/E0gouOpBhoYjAArADMdmBAvQkOcBAAAAAAAAAAAAAAAAAAAAAAAAAVAEDAAAAAA41TPIKEAkUbhdWZzBAA+AQCAQAAv7LdOZGZe90Di6CAAAQLAAAAAAQAAAAAAAAAAAAAAAAAAAAAABcgAkEAtBQYAcGAlBwcAAAAWAw8AEDAAAAAAkUU2sBMAcVYsxGchBXZyNHAAYEAJAABA8uv05U9kpUUBUkLAAAA1AAAAAAABAAAAAAAAAAAAAAAAAAAA8wrDAwVAEGAsBAbAAHAhBAcAUGAyBwcAAAAaAwkAAAAnAw7+WIAAAQMTB1U32pr/3IH/PUgMSIQ6M6ctkGAAAAZAAAAA8BAAAALAAAA3BQaA4GAkBwbAcHAzBgLAkGAtBQbAUGAyBwcAkGA2BQZAMGAvBgbAQHAyBwbAwGAwBQYA4GAlBAbA8FAjBwdAUDAuBQMAgGAyAAdAgHA5BQZAcHA5BAAAAAAAAAAAAAAaAAAAA"
ConsoleWrite( "Base64 encoded PIDL: $sEncodedPidl = " & $sEncodedPidl & @CRLF )

$sDecodedPath = _EncryptedPIDLDecode( $sEncodedPidl )
ConsoleWrite( "Base64 decoded path: $sDecodedPath = " & $sDecodedPath & @CRLF & @CRLF )


Func _EncryptedPIDLEncode2( $sPath )

  ; Get PIDL pointer from path
  Local $pPidl = _WinAPI_ShellILCreateFromPath( $sPath )

  ; Get size of PIDL structure as the pointer is pointing to
  Local $iSize = ILGetSize( $pPidl )

  ConsoleWrite( "$pPidl = " & $pPidl & " Size is " & $iSize & @CRLF )

  ; Create a byte structure from memory location defined by PIDL pointer
  Local $tPidl = DllStructCreate( "byte[" & $iSize & "]", $pPidl )

  ; Get binary string from byte structure
  Local $sBinStr = DllStructGetData( $tPidl, 1 )

  ; Create Base64 encoded string
  Local $sEncoded = _Base64Encode( $sBinStr )

  Return $sEncoded

EndFunc


Func ILGetSize( $pidl )
  Local $aRet = DllCall( "shell32.dll", "uint", "ILGetSize", "ptr", $pidl )
  If @error Then Return SetError(1, 0, 0)
  Return $aRet[0]
EndFunc

;By trancexx: http://www.autoitscript.com/forum/index.php?showtopic=81332
Func _Base64Encode($input)

    $input = Binary($input)

    Local $struct = DllStructCreate("byte[" & BinaryLen($input) & "]")

    DllStructSetData($struct, 1, $input)

    Local $strc = DllStructCreate("int")

    Local $a_Call = DllCall("Crypt32.dll", "int", "CryptBinaryToString", _
            "ptr", DllStructGetPtr($struct), _
            "int", DllStructGetSize($struct), _
            "int", 1, _
            "ptr", 0, _
            "ptr", DllStructGetPtr($strc))

    If @error Or Not $a_Call[0] Then
        Return SetError(1, 0, "") ; error calculating the length of the buffer needed
    EndIf

    Local $a = DllStructCreate("char[" & DllStructGetData($strc, 1) & "]")

    $a_Call = DllCall("Crypt32.dll", "int", "CryptBinaryToString", _
            "ptr", DllStructGetPtr($struct), _
            "int", DllStructGetSize($struct), _
            "int", 1, _
            "ptr", DllStructGetPtr($a), _
            "ptr", DllStructGetPtr($strc))

    If @error Or Not $a_Call[0] Then
        Return SetError(2, 0, ""); error encoding
    EndIf

    Return DllStructGetData($a, 1)

EndFunc   ;==>_Base64Encode

Func _EncryptedPIDLDecode($sEncryptedPIDL)
    Local Const $CRYPT_STRING_BASE64 = 0x00000001

    ; Determine length of EncryptedPIDL
    Local $iLength = StringLen($sEncryptedPIDL)

	ConsoleWrite( "Got here $sEncodedPidl = string length " & $iLength & @CRLF )
    ; Create structure to hold the attributes returned and used during decoding
    Local $tDecodeAttribs = DllStructCreate("int Size; int Skipped;int Flags")
    If @error Then Return SetError(1,0,'')

    ; Get pointers to the attributes returned and used during decoding
    Local $pDecodeSize    = DllStructGetPtr($tDecodeAttribs, 'Size')
    Local $pDecodeSkipped = DllStructGetPtr($tDecodeAttribs, 'Skipped')
    Local $pDecodeFlags   = DllStructGetPtr($tDecodeAttribs, 'Flags')

    ; Get the values needed for decoding
    DllCall('Crypt32.dll','int','CryptStringToBinary', _
                'str',$sEncryptedPIDL, _
                'int',$iLength, _
                'int',$CRYPT_STRING_BASE64, _
                'ptr',0, _
                'ptr',$pDecodeSize, _
                'ptr',$pDecodeSkipped, _
                'ptr',$pDecodeFlags)
    ConsoleWrite( "First done " & $iLength & @CRLF )

    If @error Then
	   ConsoleWrite( "Got here 2 $sEncodedPidl a new error = " & $sEncryptedPIDL & @CRLF )
	   Return SetError(2,0,'')
	EndIf

    ; Create structure to hold the PIDL and path
    Local $tDecodeResults = DllStructCreate("char PIDL[" & DllStructGetData($tDecodeAttribs, 'Size') + 1 & "]; char Path[261]")
    If @error Then
        ConsoleWrite( "Got here 3 $sEncodedPidl a new error = " & $sEncryptedPIDL & @CRLF )
        Return SetError(3,0,'')
    EndIf
    ; Get pointers to the PIDL and  path
    Local $pDecodePIDL = DllStructGetPtr($tDecodeResults, 'PIDL')
    Local $pDecodePath = DllStructGetPtr($tDecodeResults, 'Path')

    ; Get the PIDL
    DllCall('Crypt32.dll','int','CryptStringToBinary', _
                'str',$sEncryptedPIDL, _
                'int',$iLength, _
                'int',$CRYPT_STRING_BASE64, _
                'ptr',$pDecodePIDL, _
                'ptr',$pDecodeSize, _
                'ptr',$pDecodeSkipped, _
                'ptr',$pDecodeFlags)
    If @error Then
        ConsoleWrite( "Got here 4 $sEncodedPidl a new error = " & $sEncryptedPIDL & @CRLF )
        Return SetError(4,0,'')
    EndIf

    ; Get the Path
    DllCall('Shell32.dll','int','SHGetPathFromIDList', _
                'ptr',$pDecodePIDL, _
                'ptr',$pDecodePath)
    If @error Then
        ConsoleWrite( "Got here 5 $sEncodedPidl a new error = " & $sEncryptedPIDL & @CRLF )
        Return SetError(5,0,'')
    EndIf

	ConsoleWrite( "Returnin' boss " & $iLength & @CRLF )

    ; Return the Path
    Return DllStructGetData($tDecodeResults, 'Path')
EndFunc