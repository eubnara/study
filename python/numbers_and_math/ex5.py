name = 'eub'
age = 28 # really
height = 177 # centimeters
weight = 77 # kilograms
eys = 'brown'
teeth = 'White'
hair = 'black'

print "Let's talk about %r." % name
print "He's %f inches tall." % (round(height * 0.393701))
print "He's %r kilograms heavy." % weight
print "Actually that's not too heavy."
print "He's got %s eyes and %s hair." % (eys, hair)
print "His teeth are usally %s depending on the coffe." % teeth

# this line is tricky, try to get it exactly right
print "If I add %d, %d, and %d I get %d." % (age , height, weight, age + height + height )
