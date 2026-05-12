lista = [3,2,2,3,4]

listaNumerosSemRepeticao = []
listaNumerosComRepeticao = []

controlador = False
contador = 0
listaNumeroRepeticoes = []

for i in lista:
    if lista.count(lista[i]) >1:
        if controlador == False:
            listaNumerosComRepeticao.append(i)
            controlador=True
        else:
            controlador= False
    else:  
        listaNumerosSemRepeticao.append(i)

print(listaNumerosSemRepeticao)

print(listaNumerosComRepeticao)

for i in listaNumerosComRepeticao:
    listaNumeroRepeticoes.append(lista.count(i))

print(listaNumeroRepeticoes)

print(max(listaNumeroRepeticoes))
print(lista.index(max(listaNumeroRepeticoes)))