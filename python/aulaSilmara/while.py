import random  # Importa a biblioteca para gerar números aleatórios

numero_secreto = random.randint(1, 100)  # Sorteia um número entre 1 e 100
tentativas = 0  

print("Tente adivinhar o número secreto entre 1 e 100!")

while True:
    chute = int(input("Digite seu chute: "))  
    tentativas += 1  

    if chute < numero_secreto:
        print("Chutou baixo!")  
    elif chute > numero_secreto:
        print("Chutou alto!")  
    else:
        print(f"Parabéns! Você acertou em {tentativas} tentativas!")  
        break  # Encerra o loop quando acertar