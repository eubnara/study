tabby_cat = "\tI'm tabbed in."
persian_cat = "I'm split\non a line."
backshlash_cat = "I'm \\ a \\ cat."

fat_cat = """
I'll do a list:
\t* Cat food
\t* Fishies
\t* Catnip\n\t* Grass
"""

print tabby_cat
print persian_cat
print backshlash_cat
print fat_cat

while True:
    for i in ["/", "-", "|", "\\", "|"]:
        print "%s\r" % i,
