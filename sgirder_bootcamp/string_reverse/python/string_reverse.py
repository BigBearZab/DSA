def string_rev(input:str) -> str:
    output = ""
    for l in input:
        # print(l)
        output = l + output
        # print(output)
    return output

if __name__ == "__main__":
    print(string_rev("hello"))