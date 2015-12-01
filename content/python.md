========
 python
========

In Python ist alles Syntax-Sugar

---------------------------------------
>>> l = []

>>> l = list()
---------------------------------------
>>> class AB: name="AB";

>>> AB.name
"AB"

>>> getattr(AB, "name")
"AB"

>>> type(AB).__getattribute__(AB, "name")
"AB"
---------------------------------------
>>> 2**2
4

>>> int(2).__pow__(2)
4
---------------------------------------
>>> class A(object):
  def af(self):
    print "FUNKTION a in Klasse A"

>>> a = A()
>>> a.af()
"FUNKTION a in Klasse A"
>>> type(A).__getattribute__(A, "af")(a)
"FUNKTION a in Klasse A"
---------------------------------------

>>> import types
>>> globals
<built-in function globals>
>>> globals()
{'__builtins__': <module '__builtin__' (built-in)>, '__name__': '__main__',  '__package__': None}
>>> C = type("C", (object,), dict(name="C"))
>>> globals()
{'__builtins__': <module '__builtin__' (built-in)>, '__name__': '__main__', 'C': <class '__main__.C'>, '__doc__': None, '__package__': None}
>>> 
ltins__
<module '__builtin__' (built-in)>
>>> dir(__builtins__)
['ArithmeticError', 'AssertionError', 'AttributeError', 'BaseException', 'BufferError', 'BytesWarning', 'DeprecationWarning', 'EOFError', 'Ellipsis', 'EnvironmentError', 'Exception', 'False', 'FloatingPointError', 'FutureWarning', 'GeneratorExit', 'IOError', 'ImportError', 'ImportWarning', 'IndentationError', 'IndexError', 'KeyError', 'KeyboardInterrupt', 'LookupError', 'MemoryError', 'NameError', 'None', 'NotImplemented', 'NotImplementedError', 'OSError', 'OverflowError', 'PendingDeprecationWarning', 'ReferenceError', 'RuntimeError', 'RuntimeWarning', 'StandardError', 'StopIteration', 'SyntaxError', 'SyntaxWarning', 'SystemError', 'SystemExit', 'TabError', 'True', 'TypeError', 'UnboundLocalError', 'UnicodeDecodeError', 'UnicodeEncodeError', 'UnicodeError', 'UnicodeTranslateError', 'UnicodeWarning', 'UserWarning', 'ValueError', 'Warning', 'ZeroDivisionError', '_', '__debug__', '__doc__', '__import__', '__name__', '__package__', 'abs', 'all', 'any', 'apply', 'basestring', 'bin', 'bool', 'buffer', 'bytearray', 'bytes', 'callable', 'chr', 'classmethod', 'cmp', 'coerce', 'compile', 'complex', 'copyright', 'credits', 'delattr', 'dict', 'dir', 'divmod', 'enumerate', 'eval', 'execfile', 'exit', 'file', 'filter', 'float', 'format', 'frozenset', 'getattr', 'globals', 'hasattr', 'hash', 'help', 'hex', 'id', 'input', 'int', 'intern', 'isinstance', 'issubclass', 'iter', 'len', 'license', 'list', 'locals', 'long', 'map', 'max', 'memoryview', 'min', 'next', 'object', 'oct', 'open', 'ord', 'pow', 'print', 'property', 'quit', 'range', 'raw_input', 'reduce', 'reload', 'repr', 'reversed', 'round', 'set', 'setattr', 'slice', 'sorted', 'staticmethod', 'str', 'sum', 'super', 'tuple', 'type', 'unichr', 'unicode', 'vars', 'xrange', 'zip']
>>> 

