lista_participantes = []
resultados = []
while True:
    entrada = input()
    if entrada == "":
        break
    
    participantes, linhas = list(map(int, entrada.split(" ")))
    for i in range(participantes):
        lista_participantes.append("H")
    resultadosParciais = []
    homem = 0
    rato = 0
    elefante = 0
    for i in range(linhas):
        linha = input()
        operacao,a,b = list(linha.split(" "))

        if operacao == "C":
            for x in lista_participantes[int(a)-1:int(b)]:
                if x == "H":
                    homem += 1
                elif x == "E":
                    elefante += 1
                else:
                    rato += 1
            resultadosParciais.append(f"{homem} {elefante} {rato}")
        else:
            c = int(a)
            for x in lista_participantes[int(a)-1:int(b)]:
                if x == "H":
                    lista_participantes[c-1] = "E"
                    c += 1
                elif x == "E":
                    lista_participantes[c-1] = "R"
                    c += 1

                else:
                    lista_participantes[c-1] = "H"
                    c += 1
    resultados.append(resultadosParciais)

print("\n")
for i in resultados:
    for x in i:
        print(x)
    print("\n")