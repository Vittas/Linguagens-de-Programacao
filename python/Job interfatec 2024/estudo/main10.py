sequencia = input()
sequencia = list(map(int, sequencia.split(" ")))
if all(i < j for i, j in zip(sequencia, sequencia[1:])):
    print("C")
elif all(i > j for i, j in zip(sequencia, sequencia[1:])):
    print("D")
else:
    print("N")