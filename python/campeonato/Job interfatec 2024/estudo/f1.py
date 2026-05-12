entrada = input()
entrada = list(map(int, entrada.split(" ")))
resultados = []
g, p = entrada
for _ in range(entrada[0]):
    ordem_chegada = input()
    ordem_chegada = list(map(int, ordem_chegada.split(" ")))
s = int(input())
for _ in range(s):
    pontuacao = input()
    pontuacao = list(map(int, pontuacao.split(" ")))
    pontuacao.pop(0)
    n = pontuacao.index(max(pontuacao))
    resultados.append(ordem_chegada[n])
print(resultados)