gangorra = input()
gangorra = list(map(int, gangorra.split(" ")))
peso1 = gangorra[0]*gangorra[1]
peso2 = gangorra[2]*gangorra[3]
if peso1 == peso2:
    print(0)
elif peso1 > peso2:
    print(-1)
else:
    print(1)