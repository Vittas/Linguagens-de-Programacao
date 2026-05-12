import axios from "axios";
import { createContext } from "react"

interface IProductApi {
    createProduct(nome: string, descricao: string): any
    getProduct(): any
}

export const productApiContext = createContext({} as IProductApi)

export const ProductApiProvider = ({ children }: any) => {

    const createProduct = async (nome: string, descricao: string) => {
        try {
            const resposta = await axios.post("http://localhost:4001/api/itens", {
                nome,
                descricao,
            });


            console.log("Produto criado:", resposta.data);

        } catch (erro) {
            console.error("Erro ao criar o produto:", erro);
        }

    };

    const getProduct = async () => {
        try {
            const resposta = await (await axios.get("http://localhost:4001/api/itens"))
            return(resposta.data)
            console.log(resposta)
        } catch (erro){
            console.log("Não foi possivel encontrar nenhum item:", erro);
        }
    }


    return (
        <productApiContext.Provider value={{ createProduct, getProduct }}>
            {children}
        </productApiContext.Provider>

    )
}