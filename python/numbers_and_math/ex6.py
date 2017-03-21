x = "There are %d types of people." % 10
binary = "binary"
do_not = "don't"
y = "Those who know %s and those who %s." % (binary, do_not)

print x
print y

print "I said: %r." %x # we use %r for debugging, since it displays the "raw" data of the variable, but we use %s and others for displaying to users.
print "I also said: '%s'." % y

hilarious = False
joke_evaluation = "Isn't that joke so funny?! %r" # can express formatting in a string

print joke_evaluation % hilarious

w = "This is the left side of..."
e = "a string with a right side."

print w + e
