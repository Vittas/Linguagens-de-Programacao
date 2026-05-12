lista_frase = input().upper().split()
letraigual = True
if len(lista_frase[0]) == len(lista_frase[1]):
    for letra_primeira_palavra in lista_frase[0]:
        if lista_frase[1].count(letra_primeira_palavra) < lista_frase[0].count(letra_primeira_palavra):
            letraigual = False
    if letraigual != False:
        print("S")
    else:
        print("N")
else:
    print("N")