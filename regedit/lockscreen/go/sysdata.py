from pathlib import Path
# pip install pywin32
import win32security
import os


def get_user_sid():
    desc = win32security.GetFileSecurity(
        ".", win32security.OWNER_SECURITY_INFORMATION
    )
    sid = desc.GetSecurityDescriptorOwner()

    # https://www.programcreek.com/python/example/71691/win32security.ConvertSidToStringSid
    sid = win32security.ConvertSidToStringSid(sid)
    return sid


sid = get_user_sid()
dirname = Path(
    f'C:\ProgramData\Microsoft\Windows\SystemData\{sid}\ReadOnly')
for root, dirs, files in os.walk(dirname, topdown=False):
    for name in files:
        print(os.path.join(root, name))
    for name in dirs:
        print(os.path.join(root, name))
