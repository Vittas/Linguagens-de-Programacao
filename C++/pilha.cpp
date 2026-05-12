#include <iostream>
#include <iomanip>
#include <stack>
#include <string>
#include <bits/stdc++.h>

using namespace std;

bool isPalindromo(string frase, string frase2){
    return frase == frase2;
}

int main(){
    stack<char> pilha;

    string frase = "ovo";

    for(char letra: frase){
        pilha.push(letra);
    }

    string frase2;

    while(!pilha.empty()){
        frase2 += pilha.top();
        pilha.pop();
    }

    cout << "A frase original é: " << frase << "\nA frase invertida é: " << frase2 << endl;

    int palindromo = isPalindromo(frase, frase2);

    if (palindromo) {
        cout << "É um palíndromo!" << endl;
    } 
    else {
        cout << "Não é um palíndromo." << endl;
    }

    return 0;
}