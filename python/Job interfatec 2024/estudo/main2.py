triangulo_retangulo = []
controle = True
while True:
    base = int(input("-> "))
    if base < 2 and base != 0:
        print("ERRO")
        controle = False
        break
    if base == 0:
        break
    triangulo_retangulo.append(base)
if controle:
    for i in triangulo_retangulo:
        for j in range(i):
            tamanho = j+1
            print("*"*tamanho)
        if i == triangulo_retangulo[-1]:
            break
        else:
            print("\n")