textoFinal = []
while True:
    i=0
    mensagem = ''
    codigo = input()
    if codigo == " ":
        break
    codigo = codigo.strip()
    while i<len(codigo):
        if codigo[i]==' ':
            mensagem=mensagem+' '
        else:
            i=i+1
            mensagem=mensagem+codigo[i]
        i=i+1
    texto_cifrado = mensagem
    texto_decifrado = ""
    for letra in texto_cifrado:
        if letra.isalpha():
            base = 65 if letra.isupper() else 97
            posicao = (ord(letra) - base + 13) % 26 + base
            texto_decifrado += chr(posicao)
        else:
            texto_decifrado += letra
    textoFinal.append(texto_decifrado)
for i in textoFinal:
    print(i)