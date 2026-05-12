from sympy import Matrix

tipoMatriz : str = input("Digite o tipo da matriz (2x2, 3x3, 4x4...): ")

tipoMatriz = tipoMatriz[0:1]

listaEquacao : list = []

elementoEquacaoComSimbolo : str = ''
elementoEquacao : str = ''

listaMatrizesComSimbolo : list = []

listaMatrizDeterminanteTotal : list = []
dicionarioMatriz : dict = {'+': '', '-': '', '*': '', '/': '', '=': ''}

countString = 0

for i in range(0, int(tipoMatriz)):
    equacao: str = input("Digite a equação da linha " + str(i + 1) + ": ")
    listaEquacao.append(equacao)


for equacao in range(listaEquacao.__len__()):
    listaElementoEquacao : list = []
    listaElementoEquacaoComSimbolo : list = []

    for letra in range(listaEquacao[equacao].__len__()):
            if(listaEquacao[equacao][letra]):
                if listaEquacao[equacao][letra] in dicionarioMatriz:

                    if(listaEquacao[equacao][letra] == '-'):
                        listaElementoEquacao.append(elementoEquacao)
                        elementoEquacao = ''
                        elementoEquacao  += listaEquacao[equacao][letra]
                    else:
                        listaElementoEquacao.append(elementoEquacao)
                        elementoEquacao = ''
                else:
                    if(listaEquacao[equacao][letra].isnumeric()):
                        elementoEquacao += listaEquacao[equacao][letra]
                        
                        if listaEquacao[equacao].__len__() - 1 == letra:
                            elementoEquacao = ''

    listaElementoEquacao = list(filter(None, listaElementoEquacao))
    listaMatrizesComSimbolo.append(listaElementoEquacaoComSimbolo)
    listaMatrizDeterminanteTotal.append(listaElementoEquacao)


print(listaMatrizDeterminanteTotal)
print(listaMatrizesComSimbolo)


DeterminanteMatriz = Matrix(listaMatrizDeterminanteTotal)

determinanteTotal = DeterminanteMatriz.det()

print(determinanteTotal)

# for (elemento in range)
            # if listaEquacao[i][equacao][letra].isnumeric():
            #     print(listaEquacao[i][equacao][letra])
# 1x+2y+1z=7
# 2x+3y-1z=-1
# 4x-1y+2z=18