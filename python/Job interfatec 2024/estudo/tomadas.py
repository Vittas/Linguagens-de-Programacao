tomadas = input()
tomadas = tomadas.split(" ")
total = 0
for i in range(4):
    total += int(tomadas[i])

print(total-3)