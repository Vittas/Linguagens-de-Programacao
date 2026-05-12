#include <iostream>
// #include <windows.h>
#include <vector>
#include <ctime>
 
using namespace std;
 
void bubbleSort(vector<int> &numeros) {
    int tamanhoLista = numeros.size();
 
    for(int i = 0; i < tamanhoLista - 1; i++) {
        for(int j = 0; j < tamanhoLista - i - 1; j++) {
            if(numeros[j] > numeros[j + 1]){
                swap(numeros[j], numeros[j + 1]);
            }
        }
    }
}
 
int main() {
    // SetConsoleOutputCP(CP_UTF8);
    srand(time(0));
 
    vector<int> numeros;
 
    for(int i = 0; i < 9; i++){
        numeros.push_back(rand());
    }
 
    cout << "Sem ordenação: ";
    
    for(int i = 0; i < 9; i++){
        cout << numeros[i];
        if (i != 9){
            cout << " -> ";
        }
    }
    cout << "\n \n";
   
    bubbleSort(numeros);  
   
    cout << "Com ordenação: ";
    for(int i = 0; i < 9; i++){
        cout << numeros[i];
        if (i != 9){
            cout << " -> ";
        }
    }
    cout << "\n";

  return 0;
}