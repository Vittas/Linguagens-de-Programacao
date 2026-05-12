import { login } from "./login"

describe('Login', () => {
    it('Deve exibir um alert escrito "Bem vindo!" ',  () => {
        const mockAlert = jest.fn()
        window.alert = mockAlert

        login()
        expect(mockAlert).toHaveBeenCalledWith('Bem vindo!')

    })
 })