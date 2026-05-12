import random
numeroSorteado = random.randint(1, 100)
while True:
    numeroDoUsuario = int(input("digite seu numero: "))
    if numeroSorteado == numeroDoUsuario:
        print("Acertou")
        break
    elif numeroDoUsuario > numeroSorteado:
        print("O numero dito é maior que o sorteado.")
    else:
        print("O numero dito é menor que o sorteado.")