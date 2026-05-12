posicaoMuleke = list(map(int, input().split(" ")))
posicaoSaida = list(map(int, input().split(" ")))
posicaobloqueada = []
posicaobloqueada.append([posicaoMuleke[0]-1, posicaoMuleke[1]])
posicaobloqueada.append([posicaoMuleke[0]+1, posicaoMuleke[1]])
posicaobloqueada.append([posicaoMuleke[0], posicaoMuleke[1]-1])
posicaobloqueada.append([posicaoMuleke[0], posicaoMuleke[1]+1])
posicaobloqueada.append([posicaoSaida[0]-1, posicaoSaida[1]])
posicaobloqueada.append([posicaoSaida[0]+1, posicaoSaida[1]])
posicaobloqueada.append([posicaoSaida[0], posicaoSaida[1]-1])
posicaobloqueada.append([posicaoSaida[0], posicaoSaida[1]+1])
quantidadeMonstro = int(input())
posicaoMonstro = []
controle = 0
for _ in range(quantidadeMonstro):
    posicao = list(map(int, input().split(" ")))
    posicaoMonstro.append(posicao)
for i in posicaoMonstro:
        if i in posicaobloqueada:
            controle += 1
        else:
            break
if controle >= 4:
    print("-1")
else:
    passos1 = (posicaoMuleke[0] - posicaoSaida[0])
    passos2 = (posicaoMuleke[1] - posicaoSaida[1])
    if passos1< 0:
        passos1 *= -1
    if passos2 < 0:
        passos2 *= -1
    saida = passos1 + passos2
    print(saida)