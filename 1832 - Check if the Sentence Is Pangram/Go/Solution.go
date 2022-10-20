func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false 
    }
    
    alpha := make(map[rune]bool)
    for _, c := range sentence {
        if alpha[c] == false {
            alpha[c] = true  
        }
          
        if (len(alpha) == 26) {
            return true; 
        }
    }
    return false;
}