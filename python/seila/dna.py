dna = "AAGCTA"
dna = dna.upper()
dna_lenght = len(dna)

#print (dna_lenght)

result = ""

for i in dna:
    if i == "A":
        result += "T"
    elif i == "G":
        result += "C"
    elif i == "C":
        result += "G"
    elif i == "T":
        result += "A"


print(result)

n = 0

for i in dna[:: -1]:
    if result[n] != i:
        print(f"{result[n]} != {i}")

    else:
        print(f"{result[n], n + 1} = {i}")

    n += 1


# if result == dna[:: -1]:
#     print(dna[:: -1])
#     print("correto")
# elif result != dna[:: -1]:
#     for i in dna[:: -1]:
#         print(i)

#if result == dna[:: -1]:
#    print("true")
#else:
#    print("false")