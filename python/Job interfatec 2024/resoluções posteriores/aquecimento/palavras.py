texto = input().lower().split()
palavra1 = list(texto[0])
palavra2 = list(texto[1])
for i in range(len(palavra1)):
    try:
        palavra2.remove(palavra1[i])
    except ValueError:
        pass
if len(palavra2) == 0:
    print("S")
else:
    print("N")