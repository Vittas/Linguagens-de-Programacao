equa =  input("Digite uma equação: ")

num = 1
num2= 2

for i in equa:

    sinais = i



    # num =(num + num2)
    # print(num)

    
    if sinais == "+" or sinais == "-":
        num2 = i
        num2 = num2.replace("+", "0") 
        num2= int(num2)
        print(num2)
z
    # num =  num + num2
 