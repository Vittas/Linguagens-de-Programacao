informaçoes = list(map(int, input().split(" ")))
x, y = informaçoes[3] - 1, informaçoes[4]- 1
movimentos = list(input())
for i in reversed(movimentos):
    if i == "B":
        x -= 1
    elif i == "C":
        x += 1
    elif i == "D":
        y -= 1
    elif i == "E":
        y += 1
if x < informaçoes[0] or y < informaçoes[1]:
    print(-1 -1)
else:
    print(x + 1, y + 1)