import scala.io.Source


/*
 
 1. get lines into memory
 2. calculate maximum length of the string of length
 3. print line numbers with pipe bar & contents
 */

def widthOfLength (line: String) = line.length.toString.length

if (args.length > 0) {
    val lines = Source.fromFile(args(0)).getLines().toList
    val longestLine = lines.reduceLeft(
        (a, b) => if (a.length > b.length) a else b     
    )
    val maxWidth = widthOfLength(longestLine)
    for (line <- lines) {
      val numSpaces = maxWidth - widthOfLength(line)
      val padding = " " * numSpaces
      println(padding + line.length + " | " + line)
    }
}
else
    Console.err.println("Please enter filename")
