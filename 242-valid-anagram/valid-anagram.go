func isAnagram(s string, t string) bool {
    sCount:=make(map[rune]int)
    tCount:=make(map[rune]int)
    if len(s)!=len(t){
        return false
    }

     for _,ch:=range s{
        sCount[ch]++
     }
     for _,ch:=range t{
        tCount[ch]++
     }
     for ch,count:=range sCount{
      if tCount[ch]!=count{
        return false
      }
     }
    return true
}