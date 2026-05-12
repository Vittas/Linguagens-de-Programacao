lista_de_contatos = []
while(True):
    control = False
    nome = input("Adicione o nome do contato: ")

    numero = input("Adicione o número do contato: ")

    email = input("Adicione o email do contato: ")

    contato = [ nome, numero, email]

    lista_de_contatos.append(contato)

    resposta = input("Deseja visualizar algum contato específico? S/N: ").upper()
    
    if(resposta == 'S'):
        nomeBusca = input("Digite um nome: ")
        for i in lista_de_contatos:
            for nome_contato in i[0]:
                if nome_contato == nomeBusca:
                    print(i)
                    control = True


        if(control != True):
            print("Contato não encontrado")

    else:
        print(lista_de_contatos)
        break

