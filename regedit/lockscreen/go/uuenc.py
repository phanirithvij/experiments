#--*
from codecs import decode, encode
import binascii

orig = "Cat"
lockfolder = r"C:\Users\Rithvij\Desktop\Locksreen"
lockscreen_pdi = "/E"+"AFA8BUg/E0gouOpBhoYjAArADMd"+"qDAuAoOM/LtssNTCBbK/lumHacQmAQAAYCAv7bEAAAAk18O/SrbVHwKznpLy7p1BctfDAcBfadAUAw7AEDAAAAAA4JUztGEAw0bjt2cyVWZuBARAkAAEAw7+mGUSCnSRBmYuAAAAU3dSAAAAsDAAAAAAAAAAAAAAAAAAAQ/5LEAMBwbAMGArBwcAIHAlBQZA4GAAAAGAMJAAAwJA8uvFCAAAEzUQN1td66/Nyx/DFIjECkOjOXLpBAAAQGAAAAAfAAAAwCAAAwdAkGAuBAZA8GA3BwcA4CApBQbA0GAlBgcAMHApBgdAUGAjBwbA4GA0BgcA8GAsBAcAEGAuBQZAwGAfBwYAcHA1AgbAEDAoBgMAQHA4BQeAUGA3BQeAAAAAAAAAAAAAAAGAAAA"

wallps = r"D:\Images\Wallpapers"
text = "2F"+"AFA8BUg/E0gouOpBhoYjAArADMd"+"mBAvQkOcBAAAAAAAAAAAAAAAAAAAAAAAAAVAEDAAAAAA41TPIKEAkUbhdWZzBAA+AQCAQAAv7LdOZGZe90Di6CAAAQLAAAAAAQAAAAAAAAAAAAAAAAAAAAAABcgAkEAtBQYAcGAlBwcAAAAWAw8AEDAAAAAAkUU2sBMAcVYsxGchBXZyNHAAYEAJAABA8uv05U9kpUUBUkLAAAA1AAAAAAABAAAAAAAAAAAAAAAAAAAA8wrDAwVAEGAsBAbAAHAhBAcAUGAyBwcAAAAaAwkAAAAnAw7+WIAAAQMTB1U32pr/3IH/PUgMSIQ6M6ctkGAAAAZAAAAA8BAAAALAAAA3BQaA4GAkBwbAcHAzBgLAkGAtBQbAUGAyBwcAkGA2BQZAMGAvBgbAQHAyBwbAwGAwBQYA4GAlBAbA8FAjBwdAUDAuBQMAgGAyAAdAgHA5BQZAcHA5BAAAAAAAAAAAAAAaAAAAA"
print(text)

dataz = binascii.a2b_base64(text)
print(dataz)

x = binascii.b2a_uu(b'\x9a\x19\xa1b-G>T\x07=[\xd0\x88\x83\xf9*\x18R')
print(x)
# uu.encode('uuencode_decoded.dat', 'help.enc.txt')
# uu.decode('help.enc.txt', '-')
# uu.decode('data.txt', '-')
et = 'uu'
enc_data = encode(orig.encode(), et)
un_enc_data = decode(enc_data, et)
print("\n\nEncoding\t: {}".format(et))
print("Orig\t\t: {}".format(orig))
print("Encoded\t\t: {}".format(enc_data))
print("byte UnEncoded\t: {}".format(un_enc_data))
print("utf8 UnEncoded\t: {}".format(un_enc_data.decode()))
