//Dio Bank

//Abstração  name, AccountNumber

//FUNCIONALIDADES
//Depositar | Sacar

import {peopleAccount} from './class/peopleAccount';
import {companyAccount} from './class/companyAccount';


const newPeopleAccount : peopleAccount = new peopleAccount( 2, "Felipe", 400)
console.log(newPeopleAccount)
// newPeopleAccount.set_name("Vitor")
newPeopleAccount.deposit()
// console.log(newPeopleAccount)

// const company : companyAccount = new companyAccount("Embraer",124)
// console.log(company)
// company.get_Loan()
// company.deposit()