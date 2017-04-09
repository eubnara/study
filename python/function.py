# -*- coding:utf-8 -*-
def say_myself(name, old, man=True):
    print("my name is %s " % name)
    print("I'm %d years old." % old)
    if man:
        print("Male")
    else:
        print("Female")

say_myself('이유비', 28)
say_myself('이유비', 28, True)
