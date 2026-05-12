listaFinal = []
while True:
    texto = input()
    if texto == "*":
        break
    separado = texto.split()
    primeiraletra = separado [0] [0].lower()
    if all(palavra [0].lower() == primeiraletra for palavra in separado):
        listaFinal.append("Y")
    else:
        listaFinal.append("N")
for i in listaFinal:
    print(i.upper())