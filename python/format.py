data = {'first': 'Hardor', 'last' : 'Hador!'}
person = {'a': 'eub', 'b': 'ldm'}
print '%(first)s %(last)s' % data
print '{first} {last}'.format(**data)
print '{d[first]} {d[last]} {p[a]} {p[b]}'.format(d=data, p=person)
