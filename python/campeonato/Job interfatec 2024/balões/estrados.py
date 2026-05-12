vezes = int(input())
for _ in range(vezes):
    a, b, c = list(map(int, input().split(" ")))
    tamanhoRipas = b*c
    tamanhoRestante = a-tamanhoRipas
    quantidadeDeEspacoEntreAsRipas = b-1
    espaco = tamanhoRestante/quantidadeDeEspacoEntreAsRipas

    if espaco < 10:
        print("projeto superdimensionado")
    elif espaco > 20:
        print("projeto subdimensionado")
    elif espaco >= 10 and espaco <= 20:
        print("projeto ok")