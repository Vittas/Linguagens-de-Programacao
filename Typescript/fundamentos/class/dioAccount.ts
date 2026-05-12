export abstract class dioAccount{
    private name: string
    private readonly accountNumber: number
    balance: number = 0
    private status: boolean = true


    constructor(name:string, accountNumber:number){
        this.name = name
        this.accountNumber = accountNumber
    }

    set_name = (name:string):void => {
        this.name = name
        console.log("Name changed!")
    }

    get_name = ():string => {
        return this.name
    }

    deposit = ():void => {
        if(this.validate_status()){
            console.log("You deposited")
        }
    }

    withdraw = ():void =>{
        console.log("You withdrawn")
    }

    get_value = ():void => {
        console.log(this.balance)
    }

    private validate_status = (): boolean => {
        if (this.status){
            return this.status
        }
        throw new Error()
    }

}

