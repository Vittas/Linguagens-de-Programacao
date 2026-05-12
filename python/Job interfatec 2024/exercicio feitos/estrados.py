listaFinal = []
vezes = int(input())
for _ in range(vezes):
    a, b, c = list(map(int, input().split(" ")))
    tamanhoRipas = b*c
    tamanhoRestante = a-tamanhoRipas
    quantidadeDeEspacoEntreAsRipas = b-1
    espaco = tamanhoRestante/quantidadeDeEspacoEntreAsRipas

    if espaco < 10:
        listaFinal.append("projeto superdimensionado")
    elif espaco > 20:
        listaFinal.append("projeto subdimensionado")
    elif espaco >= 10 and espaco <= 20:
        listaFinal.append("projeto ok")
for i in listaFinal:
    