// assumes ASCII lowercase a-z only
func isAnagram(s, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var freq [26]uint16

    for i := range s {
        freq[s[i]-'a']++
    }

    for i := range t {
        pos := t[i] - 'a'

        if freq[pos] == 0 {
            return false
        }

        freq[pos]--
    }

    return true
}