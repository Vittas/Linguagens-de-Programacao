posicaotriancgulo = {1:"A",2:"B",3:"C",4:"D",5:"E",6:"F",7:"G",8:"H",9:"I", 0:"I"}

entrada = int(input())
x = (entrada//9)+1
z = entrada%9
if z == 0:
    x = x-1

if 1 <= entrada <= 1000000:
    print(f"{x}{posicaotriancgulo[z]}")