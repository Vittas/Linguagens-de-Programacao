contador = 1
while True:
    n = int(input())
    if n == -1:
        break
    resultado = (2**n + 1) * (2**n + 1)
    print(f"\nTeste {contador}\n{resultado}")
    contador += 1