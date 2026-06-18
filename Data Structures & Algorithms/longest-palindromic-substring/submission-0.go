func longestPalindrome(s string) string {
    result := ""

        for i:=0;i<len(s);i++{
               odd :=  expand(s,i,i)
                       even := expand(s,i,i+1)

                               if len(odd)>len(result){
                                           result = odd
                                                   }
                                                           if len(even)>len(result){
                                                                       result = even
                                                                               }
                                                                                   }

                                                                                       return result
                                                                                       }




                                                                                       func expand(s string , l, r int )string {
                                                                                           for l>=0 && r <len(s) && s[l]==s[r]{
                                                                                                   l--
                                                                                                           r++
                                                                                                               }

                                                                                                                   return s[l+1 : r]
                                                                                                                   }

