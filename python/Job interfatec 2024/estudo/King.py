resultados = []
while True:
    mao = input()
    print(mao)
    mao = mao.split(" ")
    mao.pop()
    if mao[0] == "0":
        break
    mao = list(map(int, mao))

    if all(i == mao[0] for i in mao):
        if mao[0]== 13:
            resultados.append("*")
        else:
            resultados.append(f"{mao[0]+1} {mao[1]+1} {mao[2]+1}")
    elif mao[0] == mao[1] or mao[2] == mao[0] or mao[1] == mao[2]:
        if mao[0] == mao[1] == 13 or mao[2] == mao[0] == 13 or mao[1] == mao[2] == 13:
            resultados.append(f"1 1 1")
        elif mao[0] > mao[1] < mao[2]:
            resultados.append(f"{mao[0]} {mao[1]+1} {mao[2]}")
        elif mao[0] < mao[1] == mao[2]:
            resultados.append(f"{mao[2]} {mao[1]} {mao[0]+2}")
        else:
            if mao[0] == 13:
                resultados.append(f"1 {mao[1]+1} {mao[2]+1}")
            elif mao[1] == 13:
                resultados.append(f"1 {mao[0] + 1} {mao[2]+1}")

            elif mao[2] == 13:
                resultados.append(f"1 {mao[0] + 1} {mao[1]+1} ")
            else:
                resultados.append(f"{mao[0]} {mao[1]} {mao[2]+1}")

    else:
        resultados.append("1 1 2")    
for i in resultados:

    print(i)