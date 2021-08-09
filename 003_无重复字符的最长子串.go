func lengthOfLongestSubstring(s string) int {
    var maxlen int
    substr := ""
    for _, ch := range s {
        if idx := strings.Index(substr, string(ch)); idx >= 0 {
            if len(substr) > maxlen {
                maxlen = len(substr)
            }
            // 使用slice分片截取字符串
            substr = string([]byte(substr)[idx+1:])
            substr += string(ch)
        } else {
            substr += string(ch)
        }
    }
    if len(substr) > maxlen {
        maxlen = len(substr)
    }
    return maxlen
}
