from math import sqrt
lado_triangulo = input()
listafinal = list(map(float,lado_triangulo.split(" ")))
h = max(listafinal)/2
c = max(listafinal,2)

if 0 < listafinal[0] <= 10**6  and 0 < listafinal[1] <= 10**6 and 0 < listafinal[2] <= 10**6:
    
    altura = sqrt((c*c)-(h*h))
    print(altura)
    total = (h*altura)/2
    
    print(f"{total:.2f}")