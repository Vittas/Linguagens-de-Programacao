strs = ["dog","racecar","car"]
controleInt = 0
prefix = ""
while True:
    try:
        letra = strs[0][controleInt]
        for palavra in strs:
            if palavra[controleInt] != letra:
                print(prefix)
                exit()
        prefix += letra
        controleInt += 1
    except IndexError:
        print(prefix)
        break

