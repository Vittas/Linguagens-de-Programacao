def calc():   
    num = float(input("Digite o Valor de um número: "))
    num2 = float(input("Digite o Valor de outro um número: "))

    op = input("""Digite o operador desejado: 
                [+] - Soma
                [-] - Subtração
                [*] - Multiplicação
                [/] - Divisão
                
                Digite o número do departamento desejado: """)

    match op:
        case "+":
            resultado = num + num2
            print("O resultado é: ",resultado)

        case "-":
            resultado = num - num2
            print("O resultado é: ",resultado)

        case "*":
            resultado = num * num2
            print("O resultado é: ",resultado)


        case "/":
            resultado = num / num2
            print("O resultado é: ",resultado)
            
calc()

