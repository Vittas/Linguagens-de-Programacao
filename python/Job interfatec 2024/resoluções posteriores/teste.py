def justificar_texto(largura_maxima, texto):
    palavras = texto.split()
    largura_atual = 0
    for i in range(len(palavras)):
        palavra = palavras[i]
        largura_atual += len(palavra)
        if largura_atual + len(palavras[i+1:]) + i > largura_maxima:
            espacos_necessarios = i
            largura_espaco = (largura_maxima - largura_atual) / espacos_necessarios
            print(f"{largura_espaco:.3f}", end=" ")
            largura_atual = len(palavra)
    print(0)  # Última linha alinhada à esquerda

# Leitura da entrada
largura_maxima = int(input())
texto = input()

# Chamada da função
justificar_texto(largura_maxima, texto)