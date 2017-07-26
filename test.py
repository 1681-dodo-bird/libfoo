import datetime
import ctypes


class GoInterface(ctypes.Structure):
    _fields_ = [
        ('t', ctypes.c_void_p),
        ('v', ctypes.c_void_p),
    ]

lib = ctypes.cdll.LoadLibrary('./libfoo-linux-amd64.so')
lib.FugaFuga.restype = GoInterface
lib.HogeHoge.argtypes = [GoInterface]
# lib.Fooooo()

startAt = datetime.datetime.utcnow()
lib.HogeHoge(lib.FugaFuga())
print datetime.datetime.utcnow() - startAt
