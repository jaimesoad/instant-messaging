// place files you want to import through the `$lib` alias in this folder.

import { get } from "svelte/store"
import { chatField } from "./stores"



export function ScrollToBottom() {
    const value = get(chatField)

    if (value === undefined) return
    
    setTimeout(() => {
        value.scrollTop = value.scrollHeight
    }, 100)

    chatField.set(value)
}

export function ConnectToUser() {
    
}