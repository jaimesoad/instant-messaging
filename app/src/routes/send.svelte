<script lang="ts">

    export let msg: string
</script>

<div id="wrapper">
    <dir id="main">{msg}</dir>
</div>

<style>
    #wrapper {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: end;
    }

    #main {
        max-width: 80%;
        padding: 10px;
        color: black;
        font-family: Arial, Helvetica, sans-serif;
        font-size: 1rem;
        border-radius: 1rem 1rem 0 1rem;
        background: #E0E0E0;
        margin: 0;
        word-wrap: break-word;
    }
</style>
