# coding: utf-8
import datetime
import ctypes


class GoInterface(ctypes.Structure):
    _fields_ = [
        ('t', ctypes.c_void_p),
        ('v', ctypes.c_void_p),
    ]

class GoString(ctypes.Structure):
    _fields_ = [
        ('p', ctypes.c_char_p),
        ('n', ctypes.c_int64),
    ]

    def value(self):
        return self.p[:self.n]

class Bar2_return(ctypes.Structure):
    _fields_ = [
        ('r0', GoString),
        ('r1', GoString),
    ]

# 利用する関数の定義をやる。なぜならば、.hがincludeできないからだ。
lib = ctypes.cdll.LoadLibrary('./libfoo-linux-amd64.so')
lib.FugaFuga.restype = GoInterface
lib.HogeHoge.argtypes = [GoInterface]
lib.Baaaaar.restype = GoString
lib.Bar2.restype = Bar2_return
lib.Boo.argypes = [GoString]
lib.Asdf.restype = ctypes.c_void_p
lib.Qwer.argtypes = [ctypes.c_void_p]

# overhead計測。表示される値は2回の呼び出しなので2で割った値が1回のcall
startAt = datetime.datetime.utcnow()
lib.HogeHoge(lib.FugaFuga())
print datetime.datetime.utcnow() - startAt

# goでstringを戻すケース。
ret = lib.Baaaaar()
print ret.value()

# goで複数の値を戻すケース
ret = lib.Bar2()
print ret.r0.value(), ret.r1.value()

# goにStringを渡すケース
gst = GoString()
gst.p = "hello, world from py"
gst.n = len(gst.p)
lib.Boo(gst)

# Funcをバイパス
f = lib.Asdf()
lib.Qwer(f)
