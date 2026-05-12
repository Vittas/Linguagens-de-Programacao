posicaoMonstro = []
controleMuleke = 0
controleSaida = 0
posicaoMuleke = list(map(int, input().split(" ")))
posicaoSaida = list(map(int, input().split(" ")))
posicaoBloqueadaMuleke = [[posicaoMuleke[0] - 1, posicaoMuleke[1]], [posicaoMuleke[0] + 1, posicaoMuleke[1]], [posicaoMuleke[0], posicaoMuleke[1] - 1], [posicaoMuleke[0], posicaoMuleke[1] + 1]]
posicaoBloqueadaSaida = [[posicaoSaida[0] - 1, posicaoSaida[1]], [posicaoSaida[0] + 1, posicaoSaida[1]], [posicaoSaida[0], posicaoSaida[1] - 1], [posicaoSaida[0], posicaoSaida[1] + 1]]
quantidadeMonstro = int(input())
for _ in range(quantidadeMonstro):
    posicao = list(map(int, input().split(" ")))
    posicaoMonstro.append(posicao)
controleMuleke = sum(1 for i in posicaoMonstro if i in posicaoBloqueadaMuleke)
controleSaida = sum(1 for i in posicaoMonstro if i in posicaoBloqueadaSaida)
if controleMuleke >= 4 or controleSaida >= 4:
    print("-1")
else:
    passos = abs((posicaoMuleke[0] - posicaoSaida[0])) + abs((posicaoMuleke[1] - posicaoSaida[1]))
    print(passos)