resultadoPositivo = {}
resultadoNegativo = {}
intervalos = []
for _ in range(int(input())):
    entrada = input().split(' ')
    entrada[0] = entrada[0].replace("-", "")
    entrada[1] = entrada[1].replace("-", "")
    cepInicio, cepFinal = list(map(int, entrada))
    intervalos.append([cepInicio, cepFinal])
for _ in range(int(input())):
    cepCliente = input()
    cepClienteFormatado = int(cepCliente.replace("-", ""))
    for i in intervalos:
        cepInicio = i[0]
        cepFinal = i[1]
        if cepInicio <= cepClienteFormatado <= cepFinal:
            resultadoPositivo[cepClienteFormatado] = cepCliente, "is served by our delivery system"
        elif cepInicio > cepClienteFormatado or cepFinal < cepClienteFormatado:
            resultadoNegativo[cepClienteFormatado] = cepCliente, "is not served by our delivery system"
for i in sorted(resultadoPositivo, key = resultadoPositivo.get):
    i, resultadoPositivo[i]
for i in sorted(resultadoNegativo, key = resultadoNegativo.get):
    i, resultadoNegativo[i]
for i in resultadoPositivo:
    j, k = resultadoPositivo[i]
    print(f"{j} {k}")
for i in resultadoNegativo:
    j, k = resultadoNegativo[i]
    print(f"{j} {k}")