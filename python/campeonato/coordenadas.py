lista_resultados = []
while True:

    numero_coordenada = int(input())
    if numero_coordenada == 0:
        break
    ponto_divisor = input()
    ponto_divisor = list(map(int, ponto_divisor.split(" ")))
    x , y = ponto_divisor
    for _ in range(numero_coordenada):
        ponto_coordenada = input()
        ponto_coordenada = list(map(int, ponto_coordenada.split(" ")))
        c1 , c2 = ponto_coordenada
        if c1 == x or c2 == y:
            lista_resultados.append("divisa")
        elif c1 > x and c2 > y:
            lista_resultados.append("NE")
        elif c1 < x and c2 > y:
            lista_resultados.append("NO")
        elif c1 < x and c2 < y:
            lista_resultados.append("SO")
        elif c1 > x and c2 < y:
            lista_resultados.append("SE")

for i in lista_resultados:
    print(i)