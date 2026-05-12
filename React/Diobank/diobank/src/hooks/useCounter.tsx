import { useState } from "react"

export function useCounter(initialValue: number) {

    const [count, setCount] = useState(initialValue)

    function increment() {
        setCount(initialValue + 1)
    }
    
    function decrement() {
        setCount(initialValue - 1)
    }

    return { count, increment, decrement }
}