def palindrome(s:str) -> bool:
    l = len(s) -1
    for i in range(l):
        if s[i] != s[l-i]:
            return False
    return True

if __name__ == "__main__":
    print(palindrome("aba"))