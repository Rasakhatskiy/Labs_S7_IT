<script>
  import svelteLogo from './assets/svelte.svg'
  import Counter from './lib/Counter.svelte'
  

  let text = ""
  let result = ""
  async function sendRequest(e) {
    e.preventDefault()
    const res = await (await fetch("http://127.0.0.1:1323/users/"+text)).text()
    result = res
  }
</script>

<main>
  <form>
    <input type="text" bind:value={text}>
    <button on:click={sendRequest}>Send</button>
  </form>
  <span>
    {result}
  </span>
</main>

<style>

</style>
