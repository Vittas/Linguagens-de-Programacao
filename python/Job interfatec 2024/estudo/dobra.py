quadrados_papel = 2
resposta = []
n = 1
# dobras = int(input())
# print ((quadrados_papel**dobras+1) * (quadrados_papel**dobras+1))

while True:
    dobras = int(input())
    if dobras == -1:
        break
    quadrados_papel = quadrados_papel**dobras+1
    resposta.append(quadrados_papel*quadrados_papel)
    quadrados_papel=2

for i in resposta:
    print(f"Teste {n}" )
    print(i)
    print()
    n +=1

