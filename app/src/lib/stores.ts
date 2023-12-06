import { writable } from "svelte/store";
import type { Message } from "./types";

export const messages = writable<Message[]>([])

export const chatField = writable<Element>()

export const ws = writable<WebSocket>()

export const recepient = writable<string>()
export const user = writable<string>()