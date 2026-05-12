resultadosFinais = []
dicionario_alternaticas = {0: "A" , 1:"B", 2:"C", 3:"D", 4:"E"}

while True:
    numeroQuestoes = int(input())
    if numeroQuestoes == 0:
        break
    
    for i in range(numeroQuestoes):
        linhas = input()
        linhas = linhas.split(" ")
        lista_alternativas = []
        listas_respostas = []

        contador = 0 

        for x in linhas:
            x1 = int(x)
            lista_alternativas.append(x1)

        for alternativa in lista_alternativas:
            if alternativa <= 127:
                indice = lista_alternativas.index(alternativa)
                listas_respostas.append(dicionario_alternaticas[indice])

            else:
                pass
        if len(listas_respostas) == 1:
            resultadosFinais.append(listas_respostas[0])
        else:
            resultadosFinais.append("*")            
# print(resultadosFinais)
for resultado in resultadosFinais:
    print(resultado)

    

        


        
        
        

#alternativas = [int(j) for j in linhas]