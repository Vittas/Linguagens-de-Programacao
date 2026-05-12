def calculadora_de_troco ():
    produto = input("Escolha quantos e quais produtos você irá comprar, produtos disponíveis: Feijão [R$8,99] | Arroz [R$25,60] | Salame  [R$29,99] | Alface [R$6,29] | Tomate [R$6,98]: ")



    qtd =  produto.split()


    qtd[0] = int(qtd[0])
    qtd[1] = qtd[1].lower()


    if qtd[1] == "feijão" or qtd[1] == "feijões":

        preço = 8.99

        pagamento =  preço * qtd[0]

    elif qtd[1] == "arroz":

        preço = 25.60

        pagamento =  preço * qtd[0]

    elif qtd[1] == "salame" or qtd[1] == "salames" :

        preço = 29.99

        pagamento =  preço * qtd[0]

    elif qtd[1] ==  "alface" or qtd[1] == "alfaces":

        preço = 6.29

        pagamento =  preço * qtd[0]

    elif qtd[1] == "tomate" or qtd[1] == "tomates":

        preço = 6.98

        pagamento =  preço * qtd[0]

    else:
        print("erro")

    print("O preço total  da compra deu: ",pagamento)

    dinheiro =  float(input("Informe o dinheiro para pagar: "))

    if dinheiro < pagamento:
        print("ERRO: Quantia inválida!")

    elif dinheiro >= pagamento:
        if dinheiro ==  pagamento:
            print("não há troco!")
        else:
            troco = dinheiro -  pagamento
            
            print("seu troco é de: R$ %.2f" % (troco))

calculadora_de_troco()