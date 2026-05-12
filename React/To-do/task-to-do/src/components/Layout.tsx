import { Children } from "react"
import { Footer } from "./Footer"
import { Header } from "./Header"

export const Layout = ({children}:any) =>{
    return(
        <>
            <Header/>
                <body className="grid gap-2">
                    {children}
                </body>
            <Footer/>
        </>
    )

}