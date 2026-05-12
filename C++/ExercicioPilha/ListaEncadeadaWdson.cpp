#include <iostream>

using namespace std;

struct No
{
    int data;
    No* proximo;
};

int main(){
    No* primeiroNo = new No;
    primeiroNo->data = 18;
    primeiroNo -> proximo = nullptr;

    cout << primeiroNo->data << "\n";
}