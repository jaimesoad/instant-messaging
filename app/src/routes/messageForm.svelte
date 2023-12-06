<script lang="ts">
    import { ScrollToBottom } from "$lib";
    import { messages, ws } from "$lib/stores";
    import type {Message} from '$lib/types'
    import SendIcon from "$lib/images/send.svg"

    function sendMessage(content: string) {
        const postMsg: Message = {
            sent: true,
            message: content
        }
        $messages = [...$messages, postMsg]

        $ws.send(newMessage)

        newMessage = ""

        ScrollToBottom()
    }

    function inputMessage(e: KeyboardEvent) {
        const msgTest = newMessage.trim()
        if (e.key != "Enter" || msgTest == "") return

        sendMessage(newMessage)
    }
    
    let newMessage: string = ""
</script>

<div id="main">
    <input type="text" placeholder="Write a message" on:keypress={inputMessage} bind:value={newMessage}/>
    <button on:click={() => sendMessage(newMessage)}>
        <svg>
            <image xlink:href="{SendIcon}"/>
        </svg>
    </button>
</div>

<style>
    #main {
        position: absolute;
        bottom: 0;
        width: calc(100% - 20px);
        display: inline-flex;
        padding: 0.5rem 0.625rem;
        justify-content: center;
        align-items: center;
        gap: 1.125rem;
        border-top: 1px solid #595959;
        background: #262626;
    }

    input {
        width: 52.125rem;
        height: calc(3rem - 4px);
        border-radius: 1.5rem;
        border: 2px solid #7d7d7d;
        background: #686868;
        outline: none;
        padding: 0 1.5rem;
        font-size: 1rem;
        font-family: Arial, Helvetica, sans-serif;
        color: white;
    }

    button {
        width: 3rem;
        aspect-ratio: 1/1;
        background: #4b7bf9;
        stroke-width: 2px;
        border: 2px solid #819eea;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    svg {
        width: 1.5rem;
        aspect-ratio: 1/1;
    }

    image {
        width: 1.5rem;
        aspect-ratio: 1/1;
        background: white;
    }
</style>
