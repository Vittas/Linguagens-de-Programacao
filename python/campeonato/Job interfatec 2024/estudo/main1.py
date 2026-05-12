resposta = []
while True:
    numeroQuestoes = int(input())
    if numeroQuestoes == 0:
        break
    for i in range(numeroQuestoes):
        linhas = input()
        alternativas = [int(j) for j in linhas.split()]
        contagemValidas = 0
        indiceValida = -1
        for j, k in enumerate(alternativas):
            if 0 <= k <= 127:
                contagemValidas += 1
                indiceValida = j
        if contagemValidas > 1:
            resposta.append("*")
        elif contagemValidas == 1:
            if indiceValida == 0:
                resposta.append("A")
            elif indiceValida == 1:
                resposta.append("B")
            elif indiceValida == 2:
                resposta.append("C")
            elif indiceValida == 3:
                resposta.append("D")
            elif indiceValida == 4:
                resposta.append("E")
        else:
            resposta.append("*")
for i in resposta:
    print(i)