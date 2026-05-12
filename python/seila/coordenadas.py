x = float(input("Digite um valor para a coordenada x: "))
y = float(input("Digite um valor para a coordenada y: "))

if x >= 0 and y > 0 or x > 0 and y >=0:
    print("As coordenadas caem no Quadrante 1, sendo o eixo x (",x,") e o eixo y (",y,")")
elif x <= 0 and y > 0 or x > 0 and y >=0:
    print("As coordenadas caem no Quadrante 2, sendo o eixo x (",x,") e o eixo y (",y,")")
elif x <= 0 and y < 0 or x > 0 and y <=0:
    print("As coordenadas caem no Quadrante 3, sendo o eixo x (",x,") e o eixo y (",y,")")
elif x >= 0 and y < 0 or x > 0 and y <=0:
    print("As coordenadas caem no Quadrante 4, sendo o eixo x (",x,") e o eixo y (",y,")")
elif x == 0 and y == 0 :
    print("As coordenadas caem na origem, sendo o eixo x (",x,") e o eixo y (",y,")")