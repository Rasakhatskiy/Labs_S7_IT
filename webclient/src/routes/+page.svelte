<script>
    export let data;
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<a href="/route1">route1</a>
<hr />
<a href="/route2">route2</a>

<h1>Our Databases:</h1>
<ul>
    {#each data.databases as database}
        <li><a href="/database/{database}">{database}</a></li>
    {/each}
</ul>

<style>
    /* h1 {
        background-color: blue;
    } */
</style>