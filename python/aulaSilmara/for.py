quantidadeUsuarios = int(input("Digite quantos usuarios deseja cadastrar: "))
for x in range(quantidadeUsuarios):
    nome = input("Digite um nome de usuário: ")
    senha = input("Digite uma senha: ")

    print(f'O usuário {nome} cadastrado.')

palavra = input("Digite uma palavra: ")

for letra in palavra:
    print(letra)
