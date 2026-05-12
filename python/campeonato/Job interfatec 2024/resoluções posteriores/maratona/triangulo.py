a, b, c = list(map(float, input().split(" ")))
semiPerimetro = (a+b+c)/2
area = (semiPerimetro*(semiPerimetro-a)*(semiPerimetro-b)*(semiPerimetro-c))**(1/2)
print(f"{area:.2f}")