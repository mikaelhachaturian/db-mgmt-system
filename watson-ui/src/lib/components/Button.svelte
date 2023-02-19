<script>
  export let type = "primary";
  export let flat = false;
  export let inverse = false;
</script>

<button class={type} class:flat class:inverse on:click>
  <slot />
</button>

<style>
  button {
    width: 100%;
    border: 0;
    cursor: pointer;
    border-radius: 6px;
    padding: 1em;
    font-weight: bold;
    box-shadow: 0px 8px 9px -5px rgba(0,0,0,0.38);
  }
  .primary {
    background-color: #d91b42;
    color: white;
  }
  .secondary {
    border-bottom: 5px solid #390856;
    background: rgb(116,7,222);
    background: linear-gradient(180deg, rgba(116,7,222,1) 0%, rgba(116,14,177,1) 77%, rgba(140,30,201,1) 100%);
    color: white;
  }
  .flat {
    box-shadow: none;
  }
  .primary.inverse {
    color: #d91b42;
    background: white;
    border: 2px solid #d91b42;
  }
  .secondary.inverse {
    color: #45c496;
    background: white;
    border: 2px solid #45c496;
  }
</style>
