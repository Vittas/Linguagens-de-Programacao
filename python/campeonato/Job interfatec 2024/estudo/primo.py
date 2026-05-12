number = int(input())
if number > 1:
    for i in range(2, number):
        print(i)
        if number % i == 0:
            print(number%i)
            print(number, 'não é primo')
            break
    else:
        print(number, 'é primo')