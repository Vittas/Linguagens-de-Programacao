resultado = []
while True:
    texto = input()
    if texto == "*":
        break
    texto = texto.split(" ")
    todos_corretos = True
    primeira_letra =  texto[0][0]
    for i in texto:
        if primeira_letra != i[0][0]:
            todos_corretos = False
            break
    if todos_corretos:
        resultado.append("Y")
    else:
        resultado.append("N")

for i in resultado:
    print(i)