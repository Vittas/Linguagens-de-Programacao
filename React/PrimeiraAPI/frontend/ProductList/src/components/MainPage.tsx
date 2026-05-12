import { useContext, useEffect, useState } from "react"
import { productApiContext } from "../services/ProductApiProvider";
import { CreateProductModal } from "./CreateProductModal";

export const MainPage = () => {

    const productContext = useContext(productApiContext);

    const [produtos, setProdutos] = useState<any[]>([])
    
    useEffect(()=>{
        const fetchProdutos = async () => {
            const lista = await productContext.getProduct();
            if(lista){
                setProdutos(lista)
            }
        };
        fetchProdutos();
    },[produtos]);

    return (
        <>
            <div>
                {/* <CreateProductModal/> */}
                <h2>Lista de produtos</h2>
                <ul>
                    {produtos.map((produtos, index) => (
                        <li key={index}>
                            <strong>{produtos.nome}</strong>: {produtos.descricao}
                        </li>
                    ))}
                </ul>
            </div>
        </>
    )
}