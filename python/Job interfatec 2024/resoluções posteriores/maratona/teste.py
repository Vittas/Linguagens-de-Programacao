operadores = []
elementos = []
operadorestranformados = []
expressaoOperadores = ''
expressaoElementos = ''
ordemOperadores = {'+':1,'-':2,'*':3,'/':4,'^': 5}
transformOperadores = {1:'+',2:'-',3:'*',4:'/',5:'^',}

expressao = input().upper()
elementosParenteses = expressao.split('(')

print(elementosParenteses)

for i in elementosParenteses:
    i = i.replace(")","")
    print(i)
