import scala.collection.mutable

def indexes(str: String) = {
    var result = mutable.Map[Char, mutable.Set[Int]]() 
    str.map(char => result += (char -> mutable.SortedSet[Int]()))
    0.until(str.length).map(idx => result(str.charAt(idx)) += idx)
    result
}
