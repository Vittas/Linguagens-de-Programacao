#include <iostream>
#include <iomanip>
#include <math.h>
#include <array>
using namespace std;

int main(){
    float resultado = 0;
    int nota;
    int listaDeNotas[5];
    
    for(int i = 0; i < 5; i++){
        cout << "Digite a nota do aluno " << i+1 << ": ";
        cin >> nota;
        listaDeNotas[i] = nota;
    }

    for(int i = 0; i < sizeof(listaDeNotas)/4 ;i++){
        resultado+= listaDeNotas[i];
    }

    cout << "Média: " << resultado/5;
    return 0;
}