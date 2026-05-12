for _ in range(int(input())):
    a, b, c = list(map(int, input().split(" ")))
    espaco = a-(b*c)/(b-1)
    if espaco < 10:
        print("projeto superdimensionado")
    elif espaco > 20:
        print("projeto subdimensionado")
    elif espaco >= 10 and espaco <= 20:
        print("projeto ok")