import { multiplica, soma } from "./soma"

describe('Soma', () =>{
    it('Deve somar 1 ao número informado', () => {
        const value = soma(1)
        
        expect(value).toBe(2)
    })

    it('Deve multiplicar o número informado pelo multiplicador informado', () => {
        const value = multiplica(2,3)

        expect(value).toBe(6)
    })

})

