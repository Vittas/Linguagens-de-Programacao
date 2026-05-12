lista = []
contador_de_senhas = 1 

def menu():
    print(f' 1. Gerar nova senha \n 2. Chamar próxima senha \n 3. Exibir fila de espera \n 4. Sair\n')

def GerarSenha():
    lista.append(f'P-{contador_de_senhas}')
    print("Senha adicionada")

def RemoverSenha():
    if(len(lista) == 0):
        print("A lista está vazia!")
        return 0
    print(f'Senha chamada: {lista[0]}')
    lista.remove(lista[0])


def ExibirSenha():
    # print(lista)
    for senhas in lista:
        print(senhas)

while(True):

    menu()

    op = int(input("Escolha uma das opções acima: "))

    match op:
        case 1:
            GerarSenha()
            contador_de_senhas += 1
        case 2:
            RemoverSenha()
        case 3:
            ExibirSenha()
        case 4:
            break
    

