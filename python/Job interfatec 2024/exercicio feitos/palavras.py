frase = input().upper()

frase = frase.split(" ")
letras = []
for i in frase:
    letrasPalavras = i.split()
    letras.append(letrasPalavras)
if len(frase[1]) > 30 or len(frase[0]) > 30:
  pass
elif letras[1] in letras[0]:
    print("N")
else:
    print("S")
