matriz = []
passosInversos = {"C":"B","D":"E","B":"C","E":"D"}
entrada = input()
linhas,colunas,passos,linhafesta,colunafesta = map(int,entrada.split(' '))

posicaofesta = [linhafesta,colunafesta]

for linha in range(linhas):
    for coluna in range(colunas):
        matriz.append([linha+1,coluna+1])
entrada = input()
operacoesInversas = entrada[::-1]

for op in entrada:
    opI = passosInversos[op]
    if opI == "C":
        posicaofesta[0] += 1
    elif opI == "B":
        posicaofesta[0] -= 1
    elif opI == "D":
        posicaofesta[1] += 1
    elif opI =="E":
        posicaofesta[1] -= 1
    

if posicaofesta in matriz:
    print(f"{posicaofesta[0]} {posicaofesta[1]}")
else:
    print("-1 -1")

