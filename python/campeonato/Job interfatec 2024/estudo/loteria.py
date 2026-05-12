    # inputs numero do sorteios do passado | numeros permitidos de 1 até 10 | valor total para a aposta  até 100 no maximo
while True:    
    sorteio_info = input()

    sorteio_info = sorteio_info.split(" ")

    sorteios_feitos = int(sorteio_info[0])
    tamanho_aposta = int(sorteio_info[1])
    valorlimite_aposta = int(sorteio_info[2])

    for i in range(sorteios_feitos + 1):
        aposta = input()
        