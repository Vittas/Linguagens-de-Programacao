import {dioAccount} from './dioAccount'


export class companyAccount extends dioAccount{
    constructor(name:string, accountNumber:number){
        super(name, accountNumber)
    }
    
    get_Loan = ():void => {
        console.log("You realized a loan!")
    }

    deposit = ():void => {
        console.log("The company deposited")
    }
}