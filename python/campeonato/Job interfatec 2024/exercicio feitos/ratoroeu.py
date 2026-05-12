numeroespaços = 0
largura = int(input())
frase = input().split(" ")
teste = 0
palavras = ''

for i in frase:
    palavras += i
    if (len(palavras) + numeroespaços) < largura:
        print(palavras)
        teste = numeroespaços+len(i)
    else:
        print(i)
        break
    numeroespaços += 1
