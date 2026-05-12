nomeCadastrado = "admin"
senhaCadastrada = "1234"

nome = input('digite um nome:').lower()
senha = input('digite uma senha: ')

if nome == nomeCadastrado and senha == senhaCadastrada:
    print("Login Bem-sucedido")
elif nome == nomeCadastrado and senha != senhaCadastrada:
    print("Senha incorreta!")
elif nome != nomeCadastrado and senha == senhaCadastrada:
    print("Usuário não encontrado!")
else:
    print("Credenciais inválidas!")

    