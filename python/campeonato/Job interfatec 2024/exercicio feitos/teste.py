palavra = input()
lista_frase = palavra.split()


print(all(palavra[0]== lista_frase for i in lista_frase[0][0]))