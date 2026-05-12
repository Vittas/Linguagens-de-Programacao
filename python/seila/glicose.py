gl =  int(input("Quanto de glicose você possui?"))

if gl < 75:
    print("baixa")
elif gl <= 100:
    print("normal")
elif gl <= 140:
    print("elevada")
elif gl > 140:
    print("diabete")