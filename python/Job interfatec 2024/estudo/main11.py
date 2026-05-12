listaFinal = []
pergunta = input()
numerosProibidos = input()
numerosProibidos = set(map(int, numerosProibidos.split(" ")))  # Usando um set
while True:
    chute = input()
    if chute == " ":
        break
    chute = int(chute)
    if chute in numerosProibidos:
        listaFinal.append("sim")
    else:
        listaFinal.append("não")
for i in listaFinal:
    print(i)
print(numerosProibidos)
