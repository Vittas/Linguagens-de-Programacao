ganhadores = []
contador = 1
while True:
    numeroPessoas = int(input())
    if numeroPessoas == 0:
        break
    else:
        sequencia = input()
        sequencia = list(map(int, sequencia.split(" ")))
        for i in range(1, numeroPessoas+1):
            index = i - 1
            if i == sequencia[index]:
                ganhadores.append(i)
                break
for i in ganhadores:
    print(f"\nTeste {contador}\n{i}")
    contador += 1