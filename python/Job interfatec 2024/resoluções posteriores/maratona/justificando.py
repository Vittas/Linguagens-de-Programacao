fraseFormatada = []
linhaDaFrase = []
largura = int(input())
frase = input().split(" ")
controle = largura
for i in frase:
    tamanhoPalavra = len(i)
    if tamanhoPalavra <= controle:
        linhaDaFrase.append(i)
        controle -= tamanhoPalavra+1
        # print(f"linhaDaFrase:{linhaDaFrase}")
    elif tamanhoPalavra > controle:
        fraseFormatada.append(linhaDaFrase)
        # print(f"fraseFormatada: {fraseFormatada}")
        linhaDaFrase = []
        controle = largura
        linhaDaFrase.append(i)
        # print(f"linhaDaFrase: {linhaDaFrase}")
        controle -= tamanhoPalavra+1
fraseFormatada.append(linhaDaFrase)
# print(f"fraseFormatada: {fraseFormatada}")
print(f"final: {fraseFormatada}")
for i in fraseFormatada:
    controle2 = 0
    tamanhoEspaco = 0
    for j in i:
        controle2 += len(j)
    s = largura-controle2
    n = len(i)-1
    if n == 0:
        print('0.000')
    else:    
        tamanhoEspaco = s/n
        print(f"{tamanhoEspaco:.3f}")