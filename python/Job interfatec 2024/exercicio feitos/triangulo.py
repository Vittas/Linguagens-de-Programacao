lado_triangulo = input()
a, b , c = map(float, lado_triangulo.split(" "))

if  10**6 > a > 0 and 10**6 > b > 0 and 10**6 > c > 0:
    print(round((a*b)/2, 2))