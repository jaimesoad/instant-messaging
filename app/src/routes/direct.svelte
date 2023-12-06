<script lang="ts">
    import { messages, recepient, user, ws } from "$lib/stores";
    import type { Message } from "$lib/types";

    export let name: string;

    function NewChat() {
        $messages = [];
        $recepient = name;

        if ($ws !== undefined) $ws.close()

        $ws = new WebSocket(`ws://localhost:3000/ws?user=${$user}&recepient=${$recepient}`);

        $ws.onmessage = (e) => {
            const newMsg: Message = {
                sent: false,
                message: e.data,
            };

            $messages = [...$messages, newMsg];
        };
    }
</script>

<button id="main" on:click={() => NewChat()}>
    <svg
        width="48"
        height="48"
        viewBox="0 0 48 48"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
    >
        <circle cx="24" cy="24" r="24" fill="#D9D9D9" />
        <circle cx="40" cy="40" r="8" fill="#3DAB3B" />
    </svg>
    <p>{name}</p>
</button>

<style>
    #main {
        width: 100%;
        padding: 1.5rem 1.25rem;
        display: flex;
        align-items: center;
        border: none;
        border-bottom: 2px solid #595959;
        background: #2b2b2b;
        font-size: 1.5rem;
        font-family: Arial, Helvetica, sans-serif;
        gap: 1.5rem;
        outline: none;
        margin: 0;
        color: white;
    }

    p {
        padding: 0;
        margin: 0;
    }
</style>
