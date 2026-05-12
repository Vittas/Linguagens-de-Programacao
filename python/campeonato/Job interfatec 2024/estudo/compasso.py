notas = {'W': 1.0, 'H':1/2, 'Q': 1/4, 'E':1/8, 'S':1/16, 'T':1/32, 'X':1/64}

soma=0
contador = 0
compasso_validos=[]
#compassos é uma lista
# print(compasso)
while True:
    compasso = input().upper()
    if compasso == "*":
        break
    compassos = compasso.split("/")
    for i in compassos:
        #itera pelos indices de compassos

        soma = 0
        for  letra in i:
            #anda pelas letras de um unico indice
            # print(letra)
            soma += notas[letra]

        if soma == 1:
            contador += 1
    compasso_validos.append(contador)
    contador = 0

for i in compasso_validos:
    print(i)
        

