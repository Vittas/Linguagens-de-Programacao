def senha():
    cod_usu = input("Cadastre um Código de usúario: ")
    senha = input("Cadastre uma senha: ")

    login =  input("Para logar digite o código e a senha de usúario: ")

    
    if cod_usu == login:
        login = input("Digite a senha:")
        if senha == login:
            print("Login Realizado")
        else:
            print("ERRO: Senha incorreta!")
    else:
        print("ERRO: Usúario invalido!")

senha()