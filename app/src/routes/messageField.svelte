<script lang="ts">
    import { chatField, messages } from "$lib/stores";
    import Receive from "./receive.svelte";
    import Send from "./send.svelte";
    import type { Message } from "$lib/types";
    import { ScrollToBottom } from "$lib";

    function loadReceived() {
        setTimeout(() => {
            const msg: Message = {
                sent: false,
                message: "This is an example of received message."
            }

            $messages = [...$messages, msg]

            ScrollToBottom()

        }, 100);
    }
    //loadReceived();
</script>

<div id="main" bind:this={$chatField}>
    {#each $messages as msg}
        {#if msg.sent}
            <Send bind:msg={msg.message} />
        {:else}
            <Receive bind:msg={msg.message} />
        {/if}
    {/each}
</div>

<style>
    #main {
        height: calc(100% - 4rem - 20px);
        width: calc(100% - 20px);
        padding: 10px;
        overflow-y: scroll;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }
</style>
