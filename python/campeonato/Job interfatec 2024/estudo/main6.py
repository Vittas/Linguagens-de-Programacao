# Lê o número de votos
n = int(input())

# Dicionário para armazenar os votos de cada candidato
votos = {}

# Processa cada voto
for _ in range(n):
    candidato = int(input())
    
    # Se o candidato já recebeu votos, incrementa o contador
    if candidato in votos:
        votos[candidato] += 1
    else:
        # Se for a primeira vez, inicializa com 1 voto
        votos[candidato] = 1

# Encontra o candidato com mais votos
vencedor = max(votos, key=votos.get)

# Imprime o número do candidato vencedor
print(vencedor)
