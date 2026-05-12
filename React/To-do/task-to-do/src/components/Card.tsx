interface cardContent{
    title: string
}

export const Card = ({title}:cardContent) => {
    return(
        <>
            <div>
                <h1>{title}</h1>
            </div>
        </>
    )
}