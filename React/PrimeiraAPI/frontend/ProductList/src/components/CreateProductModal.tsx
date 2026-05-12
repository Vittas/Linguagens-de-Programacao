import { useState } from "react"
import { CreateProductButton } from "./CreateProductButton"

export const CreateProductModal = () =>{
    const [productName, setProductName] = useState<string>('')
    const [productDescription, setProductDescription] = useState<string>('')

    return(
        <>
            <div>
                <h1>Create a product</h1>
                
                <label htmlFor="Product Name">Name:</label>
                <input type="text" name="productName" id="productName"
                onChange={(evt)=>{setProductName(evt.target.value)}}/>


                <label htmlFor="Product Description">Description:</label>
                <input type="text" name="productDescription" id="productDescription"
                onChange={(evt)=>{setProductDescription(evt.target.value)}}/>

                <CreateProductButton productName={productName} productDescription={productDescription}/>

            </div>
        </>
    )
}