listaFinal =[]
while True:
    operadores = []
    posfixa = []
    ordem = {'+': 1, '-': 1, '*': 2, '/': 2, '^': 3}
    expressao = input().replace(' ', '')
    if expressao == ".":
        break
    for i in expressao:
        if i.isalnum():
            posfixa.append(i)
        elif i in ordem:
            while (operadores and operadores[-1] != '(' and ordem[i] <= ordem[operadores[-1]]):
                posfixa.append(operadores.pop())
            operadores.append(i)
        elif i == '(':
            operadores.append(i)
        elif i == ')':
            while operadores and operadores[-1] != '(':
                posfixa.append(operadores.pop())
            operadores.pop()
    while operadores:
        posfixa.append(operadores.pop())
    listaFinal.append(''.join(posfixa))
for i in listaFinal:
    print(i)