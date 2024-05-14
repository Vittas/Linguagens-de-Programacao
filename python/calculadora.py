equa =  input("Digite uma equação: ")

num = 0
num2 = 0

for i in equa:
    equa = equa.replace("+", "")
    # print(equa)
    sinais = i
    if sinais == "+":
        for n in equa:
            num2 = n
            print(num2)
