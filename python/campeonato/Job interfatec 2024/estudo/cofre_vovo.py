Resultados = []
while True:
    NumeroDeposito = int(input())
    if NumeroDeposito == 0:
        break
    Diferenca = 0
    Diferencas = []
    netos = {"neto1":0,"neto2":0}
    for i in range(NumeroDeposito):
        deposito = input()
        deposito = list(map(int,deposito.split(" ")))
        netos["neto1"] += deposito[0]
        netos["neto2"] += deposito[1]
        Diferencas.append(netos["neto1"] - netos["neto2"])
    Resultados.append(Diferencas)
for i in Resultados:
    Diferenca +=1
    print(f"Teste {Diferenca}")
    for x in i:
        print(x)
