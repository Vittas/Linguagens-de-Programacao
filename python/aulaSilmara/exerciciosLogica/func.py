lista = [1,2,3,4,5]
def calular_media(lista):
    media = 0
    if len(lista) == 0:
        return 0
    
    for i in lista:
        media += i
    return media/len(lista)

print(calular_media(lista))
