import ctypes
lib = ctypes.cdll.LoadLibrary('./libfoo-linux-amd64.so')
lib.Fooooo()
lib.HogeHoge()
