diametro = int(input())
medidaCaixa = input()
medidaCaixa = medidaCaixa.split(" ")
altura = int(medidaCaixa[0])
largura = int(medidaCaixa[1])
profundidade = int(medidaCaixa[2])
if min(altura, largura, profundidade) >= diametro:
    print("S")
else:
    print("N")