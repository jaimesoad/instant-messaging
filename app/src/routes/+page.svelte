<script lang="ts">
    import Direct from "./direct.svelte";
    import MessageForm from "./messageForm.svelte";
    import MessageField from "./messageField.svelte";
    import { writable } from "svelte/store";
    import { user } from "$lib/stores";
    import { ScrollToBottom } from "$lib";

    let username = ""
    let usersWs: WebSocket

    const onlineUsers = writable<string[]>([])

    function Connect() {
        if (usersWs !== undefined) {
            console.log("websocket closed at +page.svelte")
            usersWs.close()
        }

        console.log("aaaa")
        usersWs = new WebSocket(`ws://localhost:3000/allOnline?user=${username}`)
        console.log("ssss")

        usersWs.onmessage = (e) => {
            console.log(e.data)
            $onlineUsers = <string[]>JSON.parse(e.data) ?? []
            $onlineUsers.sort()
            ScrollToBottom()
        }

        $user = username
    }
</script>

<section id="chat">
    <nav>
        <button id="new-chat">&plus;</button>
        <input type="text" bind:value={username}>
        <button on:click={() => Connect()}>Connect!</button>
    </nav>

    <div id="main-cont">
        <div id="side">
            {#each $onlineUsers as name}
                <Direct bind:name />
            {/each}
        </div>
        <div id="msgs">
            <MessageField/>
            <MessageForm />
        </div>
    </div>
</section>

<style>
    :global(body) {
        height: 100vh;
        margin: 0;
        padding: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        background: #0c0c0c;
        font-size: 1rem;
        color: white;
    }

    #chat {
        width: 85.5rem;
        min-width: 47rem;
        height: calc(100% - 2rem);
        background: #1e1e1e;
    }

    nav {
        width: calc(100% - 22px);
        height: 48px;
        border: 1px solid #595959;
        background: #333;
        padding: 0 10px;
        display: flex;
        align-items: center;
    }

    #new-chat {
        display: flex;
        height: 2rem;
        aspect-ratio: 1/1;
        padding: 0rem 0.5rem 0.1875rem 0.5rem;
        justify-content: center;
        align-items: center;
        border-radius: 0.3125rem;
        border: 2px solid #595959;
        background: #494949;
        font-size: 1.5rem;
        color: white;
    }

    #main-cont {
        height: calc(100% - 50px);
        display: flex;
        scrollbar-width: none;
        -ms-overflow-style: none;
    }

    #main-cont::-webkit-scrollbar {
        display: none;
    }

    #side {
        width: calc(19.5rem - 2px);
        height: calc(100% - 2px);
        background: #252525;
        overflow-y: scroll;
        overflow-x: hidden;
        border: 1px solid #595959;
    }

    #msgs {
        width: calc(100% - 19.5rem - 2px);
        height: calc(100% - 2px);
        min-width: 30rem;
        background: url("$lib/images/rocket-doodle.webp") 0% 0% / 384px repeat;
        border: 1px solid #595959;
        position: relative;
        overflow: hidden;
    }

    @media only screen and (max-width: 85.5rem) {
        #chat {
            width: 100vw;
            height: 100vh;
        }
    }

    @media only screen and (max-height: 52rem) {
        #chat {
            height: 100vh;
        }
    }
</style>
