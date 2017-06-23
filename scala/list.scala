val oneTwo = List(1, 2)
val threeFour = List(3, 4)
val oneTwoThreeFour = oneTwo ::: threeFour
println (oneTwoThreeFour)

// :: is right operand
// oneTwoThreeFour.::(5)
val Five = 5 :: oneTwoThreeFour
println(Five)
// error
//val appendSix = Five :: 6
//println(appendSix)

val oneTwoThree = 1 :: 2 :: 3 :: Nil
println(oneTwoThree)
