func isAnagram(s string, t string) bool {
	bucket:=make(map[byte]int)
	if len(s)!=len(t){
		return false
	}

	for i:=0; i < len(s); i++{
		bucket[s[i]]++
		bucket[t[i]]--
	}
	for _,counts:= range bucket {
		if counts!=0{
			return false

		}
	}
	return true

}
