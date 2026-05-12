verba = int(input())
listaDicionarioDados = []
listaComparacoes = []
def ordenar(e):
    return e["custoCasa"], e["pessoasCasa"]
numeroConstrucoes = int(input())
for _ in range (numeroConstrucoes):
    casaInfo = input()
    
    numCasa , custoCasa, pessoasCasa = list(map(int, casaInfo.split(" ")))
    listaDicionarioDados.append({numCasa:[custoCasa, pessoasCasa]})
    listaComparacoes.append({"numCasa":numCasa,"custoCasa":custoCasa,"pessoasCasa":pessoasCasa})

listaComparacoes.sort(ordenar())       
print(listaComparacoes)






