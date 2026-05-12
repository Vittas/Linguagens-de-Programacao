import { useContext } from "react";
import { productApiContext } from "../services/ProductApiProvider";

interface createProductInterface{
  productName: string,
  productDescription: string
}

export const CreateProductButton = ({productName, productDescription}: createProductInterface) => {
  const productContext = useContext(productApiContext);
  return (
    <>
      <button
        className='bg-blue-400 p-2'
        onClick={() => productContext.createProduct(productName, productDescription)}>
        create a product
      </button>
    </>
    

  );
}