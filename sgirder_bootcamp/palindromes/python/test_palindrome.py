from main import *

tests = {
    "aba":True,
    "abc":False
}

def testPalindrome():
    for k in tests.keys():
        assert palindrome(k) == tests[k]