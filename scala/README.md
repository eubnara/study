Functional Objects

- Basic Types and Operations
- Functional Objects
- Built-in Control Structures
- Control Abstraction


- function literal
- apply method? parenthesis? object, array?




## benefits of functional style

  - concise
  - less error-prone
  - easy to test (functions without side effects)
    - assertions and tests in chapter 14



## singleton object

  - companion object (with companion class)
  - standalone object (without companion class)

## scala application

  - singleton object with main method
  - singleton object which extends App (App trait)

## Logical operations

  - logical and (&& or &)
  - logical or (|| or |)
  - &&, || => short circuit
  - &, | => evalute all
    - how short circuiting works? => by-name parameters (section 9.5)

## BITWISE operations

  - Both &, | are logical and bitwise operations in scala.
  - shift left(<<), right(>>)
  - unsigned shift right(>>>)


## Object equality

  - eq, ne (like JAVA's ==) : reference equality (sections 11.1, 11.2)
  - equals (see chapter 30 on how to write a good `equals` method)

## Operator precedence and associativity

  - Scala decides the precedence based on the **first** character of the moethods used in operator notation.
  - exception: **assignment operators**, lowest
  - Methods which end with ":" are right associativity.

## rich wrappers

  - there are many more methods on Scala's basic types.
  - You should look at the API documentation on the rich wrapper for each basic type.

## IMMUTABLE OBJECT TRADE-OFFS

  - advantages
    1. Immutable objects are often easier to reason about than mutable ones because they do not have complex state spaces that change over time.
    2. You can pass immutable objects around quite freely, whereas you may need to make defensive copies of mutable objects before passing them to other code.
    3. There is no way for threads concurrently accessing an immutable object to corrupt its state once it has been properly constructed, because no thread can change the state of an immutable.
    4. Immutable objects make safe hash table keys.

  - disadvantages
    - Immutable objects somtimes require that a large object graph be copied, where as an update could be done in its place.

## class

  - require
  - override (ex) override def toString = )

