todasAsDatas = []
ValidaDataProva = False
ValidaDataEstudo = False
prova = input().split(" de ")
dica = input().replace("emanas", "").replace("emana", "").replace("dias", "").replace("dia", "").replace("es", "").split(" ")
estudo = input().split(" de ")
diaPorExtensoEmNumero = {"um": 1, "dois": 2, "tres": 3, "quatro": 4, "cinco": 5, "seis": 6, "sete": 7, "oito": 8, "nove": 9, "dez": 10, "onze": 11, "doze": 12, "treze": 13, "quatorze": 14, "quinze": 15, "dezesseis": 16, "dezessete": 17, "dezoito": 18, "dezenove": 19, "vinte": 20, "vinte e um": 21, "vinte e dois": 22, "vinte e três": 23, "vinte e quatro": 24, "vinte e cinco": 25, "vinte e seis": 26, "vinte e sete": 27, "vinte e oito": 28, "vinte e nove": 29, "trinta": 30, "trinta e um": 31}
mesesExtenso = {"janeiro": 1, "fevereiro": 2, "marco": 3, "abril": 4, "maio": 5, "junho": 6, "julho": 7, "agosto": 8, "setembro": 9, "outubro": 10, "novembro": 11, "dezembro": 12}
mesCurtoEmNumero = {"jan": 1, "fev": 2, "mar": 3, "abr": 4, "mai": 5, "jun": 6, "jul": 7, "ago": 8, "set": 9, "out": 10, "nov": 11, "dez": 12}
diaProva = diaPorExtensoEmNumero[prova[0]]
mesProva = mesesExtenso[prova[1]]
diaEstudo = diaPorExtensoEmNumero[estudo[0]]
mesEstudo = mesesExtenso[estudo[1]]
prova = [diaProva, mesProva]
estudo = [diaEstudo, mesEstudo]
for i in range(12):
    data = input().replace(":", "").split(" ")
    mesCurto = data[0]
    diaDoMes = int(data[1])
    mesNumero = mesCurtoEmNumero[mesCurto]
    todasAsDatas.append([mesNumero, diaDoMes])
if dica[1] == "s":
    dica = int(dica[0])*7
elif dica[1] == "m":
    dica = int(dica[0])*30
else:
    dica = int(dica[0])
for i in todasAsDatas:
    if mesProva == i[0]:
        if diaProva > i[1]:
            ValidaDataProva = True
for i in todasAsDatas:
    if mesEstudo == i[0]:
        if diaEstudo > i[1]:
            ValidaDataEstudo = True
diferencaEstudoProva = mesProva - mesEstudo
if ValidaDataEstudo == True or ValidaDataProva == True:
        print("data nao existe!")
        exit()
if diferencaEstudoProva == 0:
    if (diaProva - diaEstudo*1) < dica:
        print("olha a reprovacao chegando!")
elif diferencaEstudoProva != 0:
    if ((diaProva - diaEstudo)*(mesProva - mesEstudo)) < dica:
        print("olha a reprovacao chegando!")  
if diaEstudo == diaProva or diaEstudo > diaProva:
    print("esta de brincadeira?")
if diferencaEstudoProva == 0:
    if (diaProva - diaEstudo*1) > dica:
        print("jovem consciente!")
elif diferencaEstudoProva != 0:
    if ((diaProva - diaEstudo)*(mesProva - mesEstudo)) > dica:
        print("jovem consciente!")
if diferencaEstudoProva == 0:
    if (diaProva - diaEstudo*1) == dica:
        print("que caloura ousada!")
elif diferencaEstudoProva != 0:
    if ((diaProva - diaEstudo)*(mesProva - mesEstudo)) == dica:
        print("que caloura ousada!")