elevador = input()
elevador = elevador.split(" ")
leituras = int(elevador[0])
capacidadeElevador = int(elevador[1])
npessoas = 0
pessoas = [0]
pesoExcedido = False
for i in range(leituras):
    entradas_saidas = input()
    entradas_saidas = entradas_saidas.split(" ")
    saidas = int(entradas_saidas[0])
    entradas= int(entradas_saidas[1])
    npessoas -= saidas
    npessoas += entradas
    pessoas[0] = npessoas
    
    if pessoas[0] > capacidadeElevador:
        pesoExcedido = True
    else:
        pass

if pesoExcedido == False:
    print("N")
else:
    print("S")